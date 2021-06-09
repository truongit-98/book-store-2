package ws

import (
	"BookStore/orderbook-websocket/utils"
	"BookStore/requestbody"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

func ServeWsPublic(w http.ResponseWriter, r *http.Request) {
	wsBodyReq := requestbody.WSBodyRequest{}

	ws, err := Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("[ERROR]:", err)
		return
	}
	defer ws.Close()

	conn := &Connection{
		Send:     make(chan []byte, utils.CHANNEL_MESSAGE),
		WS:       ws,
		Channels: &sync.Map{},
	}

	_, jsonMsg, err := conn.WS.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
			log.Printf("[Error]: %v", err)
		}
	}

	err = json.Unmarshal(jsonMsg, &wsBodyReq)
	if err != nil {
		log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/orderbook.go:42")
	}

	sub := Subscription{
		conn,
		wsBodyReq.Channel,
	}

	// regis to channel
	Hub.Register <- sub
	go sub.ListenServer()
	sub.ListenClient()
}

