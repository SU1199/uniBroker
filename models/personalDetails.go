package models

type BrokerInfo struct {
	Broker   string `json:"broker"`
	Uid      string `json:"uid"`
	Password string `json:"password"`
	Pin      string `json:"pin"`
}

type Details struct {
	BrokerInfo []BrokerInfo `json:"Details"`
}
