package parser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"uniBroker/models"
)

func ParsePersonalData(location string) models.Details {
	var deets models.Details
	jsonFile, err := os.Open(location)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &deets)
	return deets
}

func ParseOrders(location string) models.Payload {
	var orders models.Payload
	jsonFile, err := os.Open(location)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &orders)
	return orders
}
