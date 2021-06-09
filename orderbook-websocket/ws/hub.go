package ws

import (
	"BookStore/orderbook-websocket/utils"
	"log"
	"strings"
	"sync"
)

var Hub = HubModel{
	Broadcast:         make(chan Message, 1000),
	Register:          make(chan Subscription, 10),
	UnRegister:        make(chan Subscription, 10),
	// channel conn
	PublicChannelsMutex: make(map[string]*sync.Map),
	PublicRWMutex:       &sync.RWMutex{},
}

func (hub *HubModel) Run() {
	log.Println("=== RUN ===")
	initChannels()

	go listenBroadcast(hub)
	go listenRegister(hub)
	go listenUnRegister(hub)
}

func listenBroadcast(hub *HubModel) {
	log.Println("=== listenBroadcast ===")
	for {
		select {
		case message := <-hub.Broadcast:
			sendAllInChannel(hub, message)
		}
	}
}

func sendAllInChannel(hub *HubModel, message Message) {
	log.Println("===== sendAllInChannel =====")
	hub.PublicRWMutex.RLock()
	defer hub.PublicRWMutex.RUnlock()

	log.Println(message.Channel, "message.Channel orderbook-websocket/ws/hub.go:73")
	if !checkChannelIsAllow(message.Channel) {
		return
	}
	connsIf, _ := hub.PublicChannelsMutex[message.Channel]
	if connsIf == nil {
		connsIf = &sync.Map{}
	}

	connsIf.Range(func(conn, isRegis interface{}) bool {
		func(innerConn *Connection, regis bool) {
			defer utils.RecoverErrorWithCallback(func() {
				connsIf.Delete(innerConn)
			})

			if innerConn != nil && regis {
				log.Println("=== START SEND === orderbook-websocket/ws/hub.go:99")
				innerConn.Send <- message.Data
				log.Println("=== END SEND === orderbook-websocket/ws/hub.go:101")
			}
		}(conn.(*Connection), isRegis.(bool))

		return true
	})
}

func listenUnRegister(hub *HubModel) {
	for {
		select {
		case unRegis := <-hub.UnRegister:
			log.Println("=== listenUnRegister ===")
			if !checkChannelIsAllow(unRegis.Channel) {
				continue
			}
			deleteConn(hub, &unRegis)
		}
	}

	log.Println("=== listenUnRegister orderbook-websocket/ws/hub.go:128 ===")
}

func listenRegister(hub *HubModel) {
	for {
		select {
		case register := <-hub.Register:
			// todo maybe bug appear here
			log.Println(register.Channel, "-- register.Channels orderbook-websocket/ws/hub.go:101")
			if !checkChannelIsAllow(register.Channel) {
				continue
			}

			addConn(hub, &register)
		}
	}
}
var (
	CHANNELS = map[string]bool{
		"WS_EVENT_PUBLIC_ORDER_BOOK": true,
		"WS_EVENT_PUBLIC_CHAT_MESSAGE": true,
	}
)

func initChannels() {
	for channel, _ := range CHANNELS {
		Hub.PublicChannelsMutex[channel] = &sync.Map{}
	}
}

func addConn(hub *HubModel, register *Subscription) {
	log.Println("add connect orderbook-websocket/ws/hub.go:172")
	hub.PublicRWMutex.Lock()
	defer hub.PublicRWMutex.Unlock()

	conns, isExist := hub.PublicChannelsMutex[register.Channel]

	// todo check more to more secure here
	if (strings.Contains(register.Channel, "WS_EVENT_UPDATE_ORDER_TRANSACTION") ||
		strings.Contains(register.Channel, "WS_EVENT_PUBLIC_CHAT_MESSAGE")) && !isExist {
		log.Println("===== is not exist =====")
		hub.PublicChannelsMutex[register.Channel] = &sync.Map{}
		conns = hub.PublicChannelsMutex[register.Channel]
	}

	//if conns == nil {
	//	log.Println("===== connection is nill =====")
	//	hub.PublicChannelsMutex[register.Channel] = &sync.Map{}
	//	conns = hub.PublicChannelsMutex[register.Channel]
	//}

	conns.Store(register.Conn, true)

	register.Conn.Channels.Store(register.Channel, true)
}

func deleteConn(hub *HubModel, unRegis *Subscription) {
	log.Println("===== deleteConn =====")
	hub.PublicRWMutex.RLock()
	defer hub.PublicRWMutex.RUnlock()

	conns := hub.PublicChannelsMutex[unRegis.Channel]
	conns.Delete(unRegis.Conn)
}

func checkChannelIsAllow(channel string) bool {
	if _, isExist := CHANNELS[channel]; isExist {
		return true
	}

	if strings.Contains(channel, "WS_EVENT_PUBLIC_ORDER_BOOK") || strings.Contains(channel, "WS_EVENT_PUBLIC_CHAT_MESSAGE") {
		return true
	}

	return false
}
