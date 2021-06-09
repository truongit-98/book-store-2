package ws

import (
	"sync"
)

//go:generate easytags $GOFILE json,xml

type Message struct {
	Data    []byte `json:"data" xml:"data"`
	Channel string `json:"channel" xml:"channel"`
}

type Subscription struct {
	Conn    *Connection `json:"conn" xml:"conn"`
	Channel string      `json:"channel" xml:"channel"`
}

type PrivateSubscription struct {
	Conn    *Connection `json:"conn" xml:"conn"`
	Channel string      `json:"channel" xml:"channel"`
	AdminID string `json:"admin_id" xml:"admin_id"`
}

type WrapChannel struct {
	Channel string            `json:"channel" xml:"channel"`
	Data    interface{}       `json:"data" xml:"data"`
	Options map[string]string `json:"options" xml:"options"`
}

// hub: maintains the set of active connections
// broadcasts: messages to the connections
type HubModel struct {
	// Registered connections
	// user here is not login yet
	// map[channel][user_connection]
	//Channels map[string]map[*Connection]bool
	PublicChannelsMutex map[string]*sync.Map `json:"public_channels_mutex" xml:"public_channels_mutex"`


	// Inbound messages from the connections
	Broadcast chan Message `json:"broadcast" xml:"broadcast"`

	// Register requests from the connections
	Register chan Subscription `json:"register" xml:"register"`

	// UnRegister requests from connections
	UnRegister chan Subscription `json:"un_register" xml:"un_register"`

	PublicRWMutex  *sync.RWMutex `json:"public_rw_mutex" xml:"public_rw_mutex"`
}
