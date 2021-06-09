package authorcontrolcontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/authorcontrolservice"
	"BookStore/services/responseservice"
	"encoding/json"
	"github.com/astaxie/beego"
	"strings"
)

type AuthorControlController struct {
	beego.Controller
}

// @Title GetControls
// @Description get controls
// @Param	token	header 	string true	"token"
// @Param	pos		query	int64	false	"Position"
// @Param	count	query	int64	false	"Count"
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *AuthorControlController) Get() {
	defer this.ServeJSON()
	pos, err := this.GetInt32("pos", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.ErrNotInt)
		return
	}
	count, err := this.GetInt32("count", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.ErrNotInt)
		return
	}
	if pos == 0 && count == 0 {
		this.GetAll()
	} else {
		this.GetPaginate(pos, count)
	}
}

func (this *AuthorControlController) GetAll() {
	roles, err := authorcontrolservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(roles)
}

func (this *AuthorControlController) GetPaginate(pos, count int32) {
	var err error
	roles, totalCount, err := authorcontrolservice.GetPaginate(pos, count)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseArray(roles, totalCount)
}

//@Title CreateNewControl
// @Description create new role by token
// @Param	token	header 	string	true	"token"
// @Param	role	body 	requestbody.AuthorControlRequest	true	"Role"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /create [post]
func (this *AuthorControlController) CreateNewControl() {
	defer this.ServeJSON()
	body := requestbody.AuthorControlRequest{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	body.Action = strings.ToUpper(body.Action)
	err = authorcontrolservice.CreateNewControl(body)
	if err != nil {
		if _, ok := responses.MapDescription[err]; ok {
			this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		} else {
			this.Data["json"] = responses.ResponseCommonSingle{
				Data: map[string]string{
					"message": err.Error(),
				},
				Error: responses.NewErr(responses.UnSuccess),
			}
		}
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponse()
}


//@Title UpdateControl
// @Description update control by token
// @Param	token	header 	string true	"token"
// @Param	body	body 	requestbody.AuthorControlPutRequest	true	"body"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /update [put]
func (this *AuthorControlController) UpdateControl() {
	defer this.ServeJSON()
	body := requestbody.AuthorControlPutRequest{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	body.Action = strings.ToUpper(body.Action)
	err = authorcontrolservice.UpdateControl(body)
	if err != nil {
		if _, ok := responses.MapDescription[err]; ok {
			this.Data["json"] = responseservice.GetCommonErrorResponse(err)
			return
		}
		this.Data["json"] = responses.ResponseCommonSingle{
			Data: map[string]string{
				"message": err.Error(),
			},
			Error: responses.NewErr(responses.UnSuccess),
		}
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponse()
}

//@Title RemoveControl
// @Description update control by token
// @Param	token	header 	string true	"token"
// @Param	controlId	path 	int32	true	"body"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /delete/:controlId [delete]
func (this *AuthorControlController) RemoveControl() {
	defer this.ServeJSON()
	controlId, err := this.GetInt32(":controlId", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = authorcontrolservice.RemoveControl(uint(controlId))
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponse()
}



