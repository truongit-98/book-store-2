package publishercontroller

import (
	_"BookStore/restapi/responses"
	"BookStore/services/publisherservice"
	"github.com/astaxie/beego"
	"BookStore/services/responseservice"

)

type PublisherController struct{
	beego.Controller
}

// @Title GetPublishers
// @Description get publishers
// @Success 200 {object} responses.ResponseCommonArray
// @router /with-book-count [get]
func (this *PublisherController) Get(){
	defer this.ServeJSON()
	data, err := publisherservice.GetAllWithBookCount()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return 
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

