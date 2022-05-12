package zerodha

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"uniBroker/models"
)

var CurrentSessionZD models.ZerodhaData

func Login(uid string, password string, pin string) error {

	type zerodha_resp struct {
		Status string `json:"status"`
		Data   struct {
			Request_id string `json:"request_id"`
		} `json:"data"`
	}

	data := url.Values{}
	data.Set("user_id", uid)
	data.Set("password", password)

	req, err := http.NewRequest("POST", "https://kite.zerodha.com/api/login", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
	}

	var t zerodha_resp
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}

	if t.Status == "success" {
		CurrentSessionZD = twoFa(uid, pin, t.Data.Request_id)
		return nil
	}
	return err
}

func twoFa(uid string, pin string, request_id string) models.ZerodhaData {
	type zerodha_resp struct {
		Status string `json:"status"`
	}

	data := url.Values{}
	data.Set("user_id", uid)
	data.Set("request_id", request_id)
	data.Set("twofa_value", pin)

	req, err := http.NewRequest("POST", "https://kite.zerodha.com/api/twofa", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var t zerodha_resp
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&t)

	if err != nil {
		log.Println(err)
	}

	var export models.ZerodhaData
	if t.Status == "success" {
		for _, cookie := range resp.Cookies() {
			if cookie.Name == "enctoken" {
				token := cookie.Value
				export.Enctoken = token
			}
		}
		export.Uid = uid
		return export
	}
	return export
}

func Margins() models.ZerodhaMargins {

	req, err := http.NewRequest("GET", "https://kite.zerodha.com/oms/user/margins", nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Authorization", "enctoken "+CurrentSessionZD.Enctoken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var t models.ZerodhaMargins
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&t)

	if err != nil {
		log.Println(err)
	}
	return t
}

func Order(inp models.ZerodhaOrder) error {

	req, err := http.NewRequest("POST", "https://kite.zerodha.com/oms/orders/"+inp.Variery, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "enctoken "+CurrentSessionZD.Enctoken)

	data := url.Values{}
	data.Set("variety", inp.Variery)
	data.Set("exchange", inp.Exchange)
	data.Set("tradingsymbol", inp.Tradingsymbol)
	data.Set("transaction_type", inp.Transaction_type)
	data.Set("order_type", inp.Order_type)
	data.Set("quantity", strconv.Itoa(inp.Quantity))
	data.Set("price", floatToString(inp.Price))
	data.Set("product", inp.Product)
	data.Set("ealidity", inp.Validity)
	data.Set("disclosed_quantity", floatToString(inp.Disclosed_quantity))
	data.Set("trigger_price", floatToString(inp.Trigger_price))
	data.Set("squareoff", floatToString(inp.Squareoff))
	data.Set("stoploss", floatToString(inp.Stoploss))
	data.Set("trailing_stoploss", floatToString(inp.Trailing_stoploss))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func floatToString(inp float64) string {
	return strconv.FormatFloat(inp, 'E', -1, 64)
}

func OrderBook() models.ZerodhaOrderBook {
	req, err := http.NewRequest("GET", "https://kite.zerodha.com/oms/orders", nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Authorization", "enctoken "+CurrentSessionZD.Enctoken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var t models.ZerodhaOrderBook
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&t)

	if err != nil {
		log.Println(err)
	}
	return t
}

func Positions() models.ZerodhaPositions {
	req, err := http.NewRequest("GET", "https://kite.zerodha.com/oms/portfolio/positions", nil)
	if err != nil {
		log.Println(err)
	}
	req.Header.Set("Authorization", "enctoken "+CurrentSessionZD.Enctoken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var t models.ZerodhaPositions
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&t)

	if err != nil {
		log.Println(err)
	}
	return t
}
