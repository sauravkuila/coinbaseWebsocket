package main

import (
	"encoding/json"
	"log"
	"net/url"
	"zerohash/business"
	"zerohash/models"

	"github.com/gorilla/websocket"
)

var websocketConn *websocket.Conn = nil

func initWebSocket(conn string) error {
	connUrl := url.URL{
		Scheme: "wss",
		Host:   conn,
	}
	log.Printf("connecting to %s", connUrl.String())

	sockConn, _, err := websocket.DefaultDialer.Dial(connUrl.String(), nil)
	if err != nil {
		log.Fatal("websocket dial error:", err)
		return err
	}
	websocketConn = sockConn
	// log.Println(websocketConn.LocalAddr().String())

	done := make(chan struct{})
	go readMessage(done)

	return nil
}

func readMessage(done chan struct{}) {
	defer close(done)
	for {
		if websocketConn != nil {
			_, message, err := websocketConn.ReadMessage()
			if err != nil {
				log.Println("websocket read error:", err)
				return
			}
			// log.Printf("recv: %s", message)
			var matchResponse models.ChannelResponse
			err = json.Unmarshal(message, &matchResponse)
			if err != nil {
				log.Printf("json unmarshall error: %s", err.Error())
			} else {
				business.CalculateVWAP(matchResponse)
			}
		}
	}
}

func writeMessage(matchRequest models.ChannelRequest) error {
	request, err := json.Marshal(matchRequest)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	err = websocketConn.WriteMessage(websocket.TextMessage, request)
	if err != nil {
		log.Println("websocket write error:", err)
		return err
	}
	return nil
}

func CloseConnection() {
	err := websocketConn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("websocket close error:", err)
		return
	}
}
