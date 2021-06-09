package categorycontroller

import (
	_"BookStore/restapi/responses"
	"BookStore/services/categoryservice"
	"github.com/astaxie/beego"
	"BookStore/services/responseservice"

)

type CategoryController struct{
	beego.Controller
}

// @Title GetCategories
// @Description get categories
// @Success 200 {object} responses.ResponseCommonArray
// @router /with-book-count [get]
func (this *CategoryController) Get(){
	defer this.ServeJSON()
	data, err := categoryservice.GetAllWithBookCount()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return 
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

