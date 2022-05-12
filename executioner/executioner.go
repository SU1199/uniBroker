package executioner

import (
	"log"
	"uniBroker/models"
	"uniBroker/zerodha"
)

func ExecuteOrders(payload models.Payload) error {
	for _, s := range payload.Payload {
		go Executioner(s)
	}
	return nil
}

func Executioner(order models.Order) error {
	if order.Broker == "zerodha" {
		var newOrder models.ZerodhaOrder
		order.Disclosed_quantity = newOrder.Disclosed_quantity
		order.Exchange = newOrder.Exchange
		order.Order_type = newOrder.Order_type
		order.Price = newOrder.Price
		order.Product = newOrder.Product
		order.Quantity = newOrder.Quantity
		order.Squareoff = newOrder.Squareoff
		order.Stoploss = newOrder.Stoploss
		order.Tradingsymbol = newOrder.Tradingsymbol
		order.Trailing_stoploss = newOrder.Trailing_stoploss
		order.Transaction_type = newOrder.Transaction_type
		order.Trigger_price = newOrder.Trigger_price
		order.Variery = newOrder.Variery
		err := zerodha.Order(newOrder)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
