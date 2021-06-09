package websocketcontroller

import (
	"BookStore/orderbook-websocket/ws"
	"github.com/astaxie/beego"
	"log"
)

type OrderbookWSController struct {
	beego.Controller
}


//@Title Get Websocket Orderbook
//@Description Get Websocket Orderbook
//@Params params body requestbody.WSBodyRequest true description
//@Success 200 {object} responses.ResponseCommonArray
//@Failure 404 nothing here
//@router / [get]
func (this *OrderbookWSController) GetPublicSocket() {
	defer this.ServeJSON()
	ws.ServeWsPublic(this.Ctx.ResponseWriter, this.Ctx.Request)

	this.Ctx.Output.SetStatus(200)

	log.Println("=== CLOSE ===")
	log.Println("=== CLOSE ===")
	log.Println("=== CLOSE ===")
}