package ws

import (
	"BookStore/requestbody"
	"BookStore/orderbook-websocket/utils"
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 30 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 30 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 1024 * 10
	OPEN           = "open"
	CLOSE          = "close"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	HandshakeTimeout: 5 * time.Second,
}

// connection: an middleman between the websocket connection and the hub
type Connection struct {
	// The websocket connection
	WS *websocket.Conn
	// Buffered channel of outbound messages
	Send chan []byte
	// todo maybe have bug here i guess
	Channels *sync.Map
}

// ReadPump pumps messages from the websocket connection to the hub.
func (sub *Subscription) ListenClient() {
	conn := sub.Conn
	defer closeClientConn(*sub, conn)

	conn.WS.SetReadLimit(maxMessageSize)
	if err := conn.WS.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:52")
		//log.Println("[ERROR]:", err)
	}
	conn.WS.SetPongHandler(func(string) error {
		if err := conn.WS.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
			log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:57")
			//log.Println("[ERROR]:", err)
		}
		return nil
	})

	listenClientMessage(conn)
}

// writePump pumps messages from the hub to the websocket connection.
func (sub *Subscription) ListenServer() {
	conn := sub.Conn
	ticker := time.NewTicker(pingPeriod)
	//defer closeClientConn(*sub, conn)
	defer ticker.Stop()

	// listen kafka event, update balance, update order....
	listenEvents(conn, ticker)
}

func (sub *PrivateSubscription) ListenServer() {
	conn := sub.Conn
	ticker := time.NewTicker(pingPeriod)
	//defer closeClientPrivateConn(*sub, conn)
	defer ticker.Stop()

	// listen kafka event, update balance, update order....
	listenEvents(conn, ticker)
}

func listenEvents(conn *Connection, ticker *time.Ticker) {
	for {
		select {
		case message, ok := <-conn.Send:
			//log.Println(string(message), "-- string(message) orderbook-websocket/ws/conn.go:100")
			if !ok {
				err := conn.write(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:106")
				}

				return
			}
			if err := conn.write(websocket.TextMessage, message); err != nil {
				if err != nil {
					log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:113")
				}
				return
			}
		case <-ticker.C:
			if err := conn.write(websocket.PingMessage, []byte{}); err != nil {
				if err != nil {
					log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:120")
				}
				return
			}
		}
	}
}

func listenClientMessage(conn *Connection) {
	log.Println("=== listenClientMessage ===")
	defer utils.RecoverError()

	for {
		_, msg, err := conn.WS.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:167")
				//log.Printf("[ERROR]: %v", err)
			}
			log.Println(err.Error(), "err.Error() orderbook-websocket/ws/conn.go:152")
			break
		}

		req := requestbody.WSBodyRequest{}
		err = json.Unmarshal(msg, &req)
		if err != nil {
			log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:130")
			continue
		}
		_, exist := conn.Channels.Load(req.Channel)

		if req.Message == OPEN {
			Hub.Register <- Subscription{
				Conn:    conn,
				Channel: req.Channel,
			}
		} else if exist && req.Message == CLOSE {
			Hub.UnRegister <- Subscription{
				Conn:    conn,
				Channel: req.Channel,
			}
		}

		log.Println("=== OKEEE ??? ===")
		//Hub.Broadcast <- Message{msg, sub.Channels}
	}
	log.Println("=== listenClientMessage END ===")
}

// write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	if err := c.WS.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:194")
		//log.Println("[ERROR]:", err)
	}
	return c.WS.WriteMessage(mt, payload)
}

func (c *Connection) writeJSON(payload interface{}) error {
	if err := c.WS.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
		log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:202")
		//log.Println("[ERROR]:", err)
	}
	return c.WS.WriteJSON(payload)
}

func closeClientConn(sub Subscription, conn *Connection) {
	log.Println("=== CLOSE  closeClientConn ===")
	// todo handle this another way
	for channel, _ := range Hub.PublicChannelsMutex {
		sub.Channel = channel
		Hub.UnRegister <- sub
	}

	close(conn.Send)
	closeWS(conn)
}


func closeWS(conn *Connection) {
	log.Println("=== closeWS === orderbook-websocket/ws/conn.go:211")
	//channels := conn.Channels
	//
	//for channel, _ := range channels {
	//	delete(Hub.Channels[channel], conn)
	//}

	if err := conn.WS.Close(); err != nil {
		log.Println(err.Error(), "-- err.Error() orderbook-websocket/ws/conn.go:216")
		//log.Println("[ERROR]:", err)
	}
}
