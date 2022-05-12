package login

import (
	"log"
	"uniBroker/models"
	"uniBroker/zerodha"
)

func Login(logins models.Details) error {
	for _, s := range logins.BrokerInfo {
		loginer(s)
	}
}

func loginer(l models.BrokerInfo) error {
	if l.Broker == "zerodha" {
		err := zerodha.Login(l.Uid, l.Password, l.Pin)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
