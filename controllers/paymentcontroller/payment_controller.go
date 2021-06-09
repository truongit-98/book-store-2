package paymentcontroller

import (
	_ "BookStore/restapi/responses"
	"BookStore/services/paymentmethodservice"
	"BookStore/services/responseservice"
	"github.com/astaxie/beego"
)

type PaymentController struct{
	beego.Controller
}

// @Title GetPayment
// @Description get Payment
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *PaymentController) Get(){
	defer this.ServeJSON()
	data, err := paymentmethodservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

