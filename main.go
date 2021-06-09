package main

import (
	"BookStore/middleware"
	"BookStore/models"
	_ "BookStore/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"log"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Connection", "Authorization", "Sec-WebSocket-Extensions", "Sec-WebSocket-Key",
			"Sec-WebSocket-Version", "Access-Control-Allow-Origin", "content-type", "Content-Type", "sessionkey", "token", "User", "Upgrade"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Content-Type", "Sec-WebSocket-Accept", "Connection", "Upgrade"},
		AllowCredentials: true,
	}))
	_, err := models.InitDB()
	if err != nil {
		log.Println(err.Error())
	}
	//models.Payment{Method: "Thanh toán khi nhận hàng (COD)"}.Create()
	//models.Payment{Method: "Thanh toán bằng thẻ ATM (Internet Banking)"}.Create()
	//models.Payment{Method: "Ví Momo"}.Create()
	//models.Payment{Method: "Ví ZaloPay"}.Create()
	//
	//(&models.Voucher{Code: "COSBEAUTY", Value: 30, Expiry: time.Now().Add(time.Duration(7*24*3600) * time.Second).Unix()}).Create()
	//(&models.Voucher{Code: "WAGMVC16H", Value: 22, Expiry: time.Now().Add(time.Duration(7*24*3600) * time.Second).Unix()}).Create()
	//(&models.Voucher{Code: "CBL2RH120KDECA", Value: 40, Expiry: time.Now().Add(time.Duration(7*24*3600) * time.Second).Unix()}).Create()
	//(&models.Voucher{Code: "FMCGCCB2", Value: 50, Expiry: time.Now().Add(time.Duration(7*24*3600) * time.Second).Unix()}).Create()


	beego.InsertFilter("/tet", beego.BeforeExec, middleware.CheckPermission())
	//go ws.Hub.Run()

	beego.Run()
}
