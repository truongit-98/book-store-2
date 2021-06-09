package bookcontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/bookservice"
	"BookStore/services/responseservice"
	"encoding/json"
	"log"
)

// @Title CreateNewBookDetail
// @Description create new book detail
// @Param token header string true "token"
// @Param body body requestbody.BookDetailPostRequestBody true "Body"
// @Success 200 {object} responses.ResponseCommonArray
// @router /create-book-detail [post]
func (this *BookController) CreateBookDetail() {
	defer this.ServeJSON()

	body := requestbody.BookDetailPostRequestBody{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		log.Println(err)
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	bType, err := bookservice.CreateBookDetail(body)
	if err != nil {
		log.Println(err)
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(bType)

}