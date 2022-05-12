package models

type Order struct {
	Timestamp          string  `json:"Timestamp"`
	Broker             string  `json:"Broker"`
	Variery            string  `json:"Variety"`
	Exchange           string  `json:"Exchange"`
	Tradingsymbol      string  `json:"Tradingsymbol"`
	Transaction_type   string  `json:"Transaction_type"`
	Order_type         string  `json:"Order_type"`
	Quantity           int     `json:"Quantity"`
	Price              float64 `json:"Price"`
	Product            string  `json:"Product"`
	Disclosed_quantity float64 `json:"Disclosed_quantity"`
	Trigger_price      float64 `json:"Trigger_price"`
	Squareoff          float64 `json:"Squareoff"`
	Stoploss           float64 `json:"Stoploss"`
	Trailing_stoploss  float64 `json:"Trailing_stoploss"`
}

type Payload struct {
	Payload []Order `json:"Payload"`
}
