package main

import (
	"log"
	"time"
	"uniBroker/upstox"
)

func main() {
	log.Println(time.Now())
	err := upstox.Login("uid", "pass", "pin")
	if err != nil {
		log.Println(err)
	}
	log.Println(upstox.CurrentSessionUpstox)
}
