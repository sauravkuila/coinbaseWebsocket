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

	initPairs := make([]string, 3)
	initPairs[0] = "ETH-EUR"
	initPairs[1] = "ETH-USD"
	initPairs[2] = "ETH-BTC"
	business.Init(initPairs, 5)

	var channel models.Channel
	channel.Name = "matches"
	channel.ProductIds = make([]string, 0)
	// channel.ProductIds = append(channel.ProductIds, "ETH-EUR")
	channel.ProductIds = append(channel.ProductIds, "ETH-USD")
	// channel.ProductIds = append(channel.ProductIds, "ETH-BTC")

	var subscribe models.ChannelRequest
	subscribe.Type = "subscribe"
	subscribe.Channels = make([]models.Channel, 0)
	subscribe.Channels = append(subscribe.Channels, channel)

	writeMessage(subscribe)
	addShutdownHook()
}

func addShutdownHook() {
	// when receive interruption from system shutdown server and scheduler
	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit

	log.Println("Interrupt called. Closing Socket Connection")
	CloseConnection()
	log.Println("Quiting application")
}
