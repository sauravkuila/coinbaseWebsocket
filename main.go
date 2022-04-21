package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"zerohash/business"
	"zerohash/models"
)

func main() {
	initErr := initWebSocket("ws-feed.exchange.coinbase.com")
	if initErr != nil {
		return
	}

	initPairs := make([]string, 0)
	initPairs = append(initPairs, "ETH-EUR")
	initPairs = append(initPairs, "ETH-USD")
	initPairs = append(initPairs, "ETH-BTC")
	business.Init(initPairs, 200)

	var channel models.Channel
	channel.Name = "matches"
	channel.ProductIds = make([]string, 0)
	channel.ProductIds = append(channel.ProductIds, initPairs...)

	var subscribe models.ChannelRequest
	subscribe.Type = "subscribe"
	subscribe.Channels = make([]models.Channel, 0)
	subscribe.Channels = append(subscribe.Channels, channel)

	writeMessage(subscribe)
	addShutdownHook()
}

func addShutdownHook() {
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	log.Println("Interrupt called. Closing Socket Connection")
	closeConnection()
	log.Println("Quiting application")
}

//sources
//https://docs.cloud.coinbase.com/exchange/docs/websocket-channels#heartbeat-channel
//https://docs.cloud.coinbase.com/exchange/reference/exchangerestapi_getaccounts
//https://docs.cloud.coinbase.com/prime/docs/websocket-feed#subscribe
