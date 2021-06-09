package commoncontroller

import (
	"BookStore/services/commonservice"
	"BookStore/services/responseservice"
	"github.com/astaxie/beego"
)

type CommonController struct {
	beego.Controller
}

// @Title GetDistrict
// @Description Get District of city
// @Param	cityId	path 	string	true	"CityID"
// @Success 200 {object} data.District
// @router /administrative/districts/:cityId [get]
func (this *CommonController) GetDistrict() {
	defer this.ServeJSON()
	cityId := this.GetString(":cityId")
	bytes, err := commonservice.LoadDistrictForCity(cityId)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(string(bytes))
}

// @Title GetWards
// @Description GetWards of district
// @Param	districtId	path 	string	true	"districtId"
// @Success 200 {object} data.District
// @router /administrative/wards/:districtId [get]
func (this *CommonController) GetWards() {
	defer this.ServeJSON()
	districtId := this.GetString(":districtId")
	bytes, err := commonservice.LoadWardsOfDistrict(districtId)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(string(bytes))
}