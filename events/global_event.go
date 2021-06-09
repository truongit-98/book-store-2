package events

import (
	"BookStore/models"
	wsmodel "BookStore/orderbook-websocket/models"
	"BookStore/orderbook-websocket/ws"

	//"BookStore/orderbook-websocket/ws"
	"log"
)

//go:generate easytags $GOFILE json,xml

type EventType int

const (
	WsOrderCreated 	EventType = iota
)
func (s EventType) String() string {
	return [...]string{"WsOrderCreated"}[s]
}

// todo delete me
var MapEventFunc = map[EventType]func(params interface{}){
	WsOrderCreated:   wsTriggerOrderCreated,
}
func BroadcastEvent(eventType EventType, params interface{}) {
	log.Println("=== " + eventType.String() + " ===")
	MapEventFunc[eventType](params)
}

type DataBody struct {
	Key   string `json:"key" xml:"key"`
	Value string `json:"value" xml:"value"`
}

type MessageBody struct {
	Body  string     `json:"body" xml:"body"`
	Data  []DataBody `json:"data" xml:"data"`
	Title string     `json:"title" xml:"title"`
}

type MetaData struct {
	Coin      string
	Currency  string
	EventType string
	Pubkey    string
	Body      map[string]string
	Title     map[string]string
}

type senders map[string]chan wsmodel.Message

func wsTriggerOrderCreated(params interface{}) {
	log.Println("method @wsTriggerTransactionCreated")

	orders := params.(*models.Order)

	ws.SendOrder2Ws(orders)
}

