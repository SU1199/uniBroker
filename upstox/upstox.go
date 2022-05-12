package upstox

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"uniBroker/models"
)

var CurrentSessionUpstox models.UpstoxData

func Login(uid, password, pin string) error {

	type upstox_resp struct {
		Success bool `json:"success"`
		Data    struct {
			TokenFor2FA string `json:"tokenFor2FA"`
		} `json:"data"`
	}

	type Data struct {
		LoginMethod string `json:"loginMethod"`
		UserID      string `json:"userId"`
		Password    string `json:"password"`
	}

	type Payload struct {
		Data Data `json:"data"`
	}

	data := Data{
		LoginMethod: "oms",
		UserID:      uid,
		Password:    base64.StdEncoding.EncodeToString([]byte(password)),
	}

	payload := Payload{Data: data}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://service.upstox.com/login/open/v2/auth/1fa?requestId=WPRO-Yf0UlHDv-9wA3DR2TBJAE", body)
	if err != nil {
		log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(resp.StatusCode)
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var t upstox_resp
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&t)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println(t)

	if t.Success {
		log.Println("2fa")

		CurrentSessionUpstox = twoFa(uid, pin, t.Data.TokenFor2FA)

		return nil
	}
	return err
}

func twoFa(uid, pin, token string) models.UpstoxData {

	type upstox_resp struct {
		Success bool `json:"success"`
		Data    struct {
			LoginKey string `json:"loginKey"`
		} `json:"data"`
	}

	type Data struct {
		TwoFAMethod      string `json:"twoFAMethod"`
		UserID           string `json:"userId"`
		InputText        string `json:"inputText"`
		TokenFor2FA      string `json:"tokenFor2FA"`
		EnableBiometrics bool   `json:"enableBiometrics"`
	}

	type Payload struct {
		Data Data `json:"data"`
	}

	data := Data{
		TwoFAMethod:      "YOB",
		UserID:           uid,
		InputText:        pin,
		TokenFor2FA:      token,
		EnableBiometrics: true,
	}

	payload := Payload{
		Data: data,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "https://service.upstox.com/login/open/v2/auth/2fa?requestId=WPRO-DUkanQbmpf1JHG2E7J63M&client_id=PW3-6Agd37PB52Q6B6DDpYWLuT7b&response_type=code&redirect_uri=https%3A%2F%2Fpro.upstox.com", body)
	if err != nil {
		log.Println(err)
	}
	req.Host = "service.upstox.com"
	req.Header.Set("Content-Length", "176")
	req.Header.Set("Sec-Ch-Ua", "\"(Not(A:Brand\";v=\"8\", \"Chromium\";v=\"101\"")
	req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device-Details", "platform=WEB|osName=Windows/10|osVersion=Chrome/101.0.4951.54|appVersion=1.3.2|manufacturer=unknown|modelName=Chrome")
	req.Header.Set("X-Client-Id", "PW3-6Agd37PB52Q6B6DDpYWLuT7b")
	req.Header.Set("X-User-Id", uid)
	req.Header.Set("Sec-Ch-Ua-Platform", "\"Linux\"")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Origin", "https://login-v2.upstox.com")
	req.Header.Set("Sec-Fetch-Site", "same-site")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Referer", "https://login-v2.upstox.com/")
	req.Header.Set("Accept-Encoding", "gzip, deflate")
	req.Header.Set("Accept-Language", "en-GB,en-US;q=0.9,en;q=0.8")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	var t upstox_resp
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&t)
	if err != nil {
		log.Println(err)
	}

	var response models.UpstoxData

	if t.Success {
		response.LoginKey = t.Data.LoginKey
		response.Uid = uid
	}

	return response

}
