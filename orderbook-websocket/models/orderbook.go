package models


//go:generate easytags $GOFILE json,xml

type Sender struct {
	Id     []string    `json:"id" xml:"id"`
	Params []string    `json:"params" xml:"params"`
	Data   interface{} `json:"data" xml:"data"`
}

type Message struct {
	Channel string      `json:"channel" xml:"channel"`
	Data    interface{} `json:"data" xml:"data"`
}

type WsWrapModel struct {
	Data   []byte `json:"data" xml:"data"`
	Pubkey string `json:"pubkey" xml:"pubkey"`
	UUID   string `json:"uuid" xml:"uuid"`
}
