package ordercontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/orderservice"
	"BookStore/services/responseservice"
	"encoding/json"
	"log"
	"github.com/astaxie/beego"
)

type OrderController struct {
	beego.Controller
}

// @Title OrderCheckout
// @Description checkout order
// @Param	token	header true	string	"token"
// @Param	body	body 	requestbody.OrderInformation	true	"OrderInfo"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /checkout [post]
func (this *OrderController) CheckoutOrder() {
	defer this.ServeJSON()
	user := this.Ctx.Request.Header.Get("User")
	if user != "" {
		body := requestbody.OrderInformation{}
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
		if err != nil {
			log.Println(err)
			this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
			return
		}
		err = orderservice.CheckoutOrder(body, user)
		if err != nil {
			this.Data["json"] = responseservice.GetCommonErrorResponse(err)
			return
		}
		this.Data["json"] = responseservice.GetCommonSucceedResponse()
		return
	}
	log.Println("user is required")
	this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
}

// @Title OrderHistory
// @Description order history
// @Param	token	header 	string true	"token"
// @Param	status 	query	int32	false	"Status"
// @Param	start 	query	int32	false	"start"
// @Param	end 	query	int32	false	"end"
// @Param	offset 	query	int32	false	"offset"
// @Param	limit 	query	int32	false	"limit"
// @Success 200 {object} responses.ResponseCommonArray
// @router /history [get]
func (this *OrderController) GetOrderHistory() {
	defer this.ServeJSON()
	status, err := this.GetInt32("status", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	start, err := this.GetInt32("start", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	end, err := this.GetInt32("end", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	offset, err := this.GetInt32("offset", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	limit, err := this.GetInt32("limit", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	data, err := orderservice.OrderHistory(status, start, end, offset, limit)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)
}