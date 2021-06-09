package utils

import "fmt"

var (
	PARAMS_SENDER   = []string{"orderbook", "ordertransaction"}
	CHANNEL_MESSAGE = 50
	CHANNEL_SEND    = 50
)

func RecoverError() {
	if r := recover(); r != nil {
		fmt.Println("panic occured: ", r)
	}
}

func RecoverErrorWithCallback(cb func())  {
	if r := recover(); r != nil {
		fmt.Println("panic occured: ", r)
		cb()
	}
}