package vouchercontroller

import (
	_ "BookStore/restapi/responses"
	"BookStore/services/responseservice"
	"BookStore/services/voucherservice"
	"github.com/astaxie/beego"
)

type VoucherController struct{
	beego.Controller
}

// @Title GetVouchers
// @Description GetVouchers
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *VoucherController) Get(){
	defer this.ServeJSON()
	data, err := voucherservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

