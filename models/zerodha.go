package models

type ZerodhaData struct {
	Enctoken   string
	Kf_session string
	Uid        string
}

type ZerodhaMargins struct {
	Status string `json:"status"`
	Data   struct {
		Equity struct {
			Enabled   bool    `json:"enabled"`
			Net       float64 `json:"net"`
			Available struct {
				Adhoc_margin    float64 `json:"adhoc_margin"`
				Cash            float64 `json:"cash"`
				Opening_balance float64 `json:"opening_balance"`
				Live_balance    float64 `json:"live_balance"`
				Collateral      float64 `json:"collateral"`
				Intraday_payin  float64 `json:"intraday_payin"`
			} `json:"available"`
			Utilised struct {
				Debts             float64 `json:"debits"`
				Exposure          float64 `json:"exposure"`
				M2m_realised      float64 `json:"m2m_realised"`
				M2m_unrealised    float64 `json:"m2m_unrealised"`
				Option_premium    float64 `json:"option_premium"`
				Payout            float64 `json:"payout"`
				Span              float64 `json:"span"`
				Holding_Sales     float64 `json:"holding_sales"`
				Turnover          float64 `json:"turnover"`
				Liquid_collateral float64 `json:"liquid_collateral"`
				Stock_collateral  float64 `json:"stock_collateral"`
				Delivery          float64 `json:"delivery"`
			}
		} `json:"equity"`
	} `json:"data"`
}

type ZerodhaOrder struct {
	Variery            string
	Exchange           string
	Tradingsymbol      string
	Transaction_type   string
	Order_type         string
	Quantity           int
	Price              float64
	Product            string
	Validity           string
	Disclosed_quantity float64
	Trigger_price      float64
	Squareoff          float64
	Stoploss           float64
	Trailing_stoploss  float64
}

type order struct {
	Placed_by                 string `json:"placed_by"`
	Order_id                  string `json:"order_id"`
	Exchange_order_id         string `json:"exchange_order_id"`
	Parent_order_id           string `json:"parent_order_id"`
	Status                    string `json:"status"`
	Status_message            string `json:"status_message"`
	Status_message_raw        string `json:"status_message_raw"`
	Order_timestamp           string `json:"order_timestamp"`
	Exchange_update_timestamp string `json:"exchange_update_timestamp"`
	Exchange_timestamp        string `json:"exchange_timestamp"`
	Variety                   string `json:"variety"`
	Modified                  string `json:"modified"`
	Exchange                  string `json:"exchange"`
	Tradingsymbol             string `json:"tradingsymbol"`
	Instrument_token          string `json:"instrument_token"`
	Order_type                string `json:"order_type"`
	Transaction_type          string `json:"transaction_type"`
	Validity                  string `json:"validity"`
	Validity_ttl              string `json:"validity_ttl"`
	Product                   string `json:"product"`
	Quantity                  string `json:"quantity"`
	Qisclosed_quantity        string `json:"disclosed_quantity"`
	Price                     string `json:"price"`
	Trigger_price             string `json:"trigger_price"`
	Average_price             string `json:"average_price"`
	Filled_quantity           string `json:"filled_quantity"`
	Pending_quantity          string `json:"pending_quantity"`
	Cancelled_quantity        string `json:"cancelled_quantity"`
	Market_protection         string `json:"market_protection"`
	Tag                       string `json:"tag"`
	Guid                      string `json:"guid"`
}

type ZerodhaOrderBook struct {
	Status string  `json:"status"`
	Data   []order `json:"data"`
}

type netPositions struct {
}
type dayPositions struct {
}

type ZerodhaPositions struct {
	Status string `json:"status"`
	Data   struct {
		Net []netPositions `json:"net"`
		Day []dayPositions `json:"day"`
	} `json:"data"`
}
