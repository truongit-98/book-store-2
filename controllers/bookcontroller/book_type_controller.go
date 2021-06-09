package bookcontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/bookservice"
	"BookStore/services/responseservice"
	"encoding/json"
	"log"
)

// @Title CreateNewBookType
// @Description create new book type
// @Param token header string true "token"
// @Param body body requestbody.BookTypePostRequestBody true "Body"
// @Success 200 {object} responses.ResponseCommonArray
// @router /create-book-type [post]
func (this *BookController) CreateBookType() {
	defer this.ServeJSON()

	body := requestbody.BookTypePostRequestBody{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		log.Println(err)
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	bType, err := bookservice.CreateBookType(body)
	if err != nil {
		log.Println(err)
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(bType)

}