package admincontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/adminservice"
	"BookStore/services/responseservice"
	"BookStore/utils"
	"encoding/json"
	"github.com/astaxie/beego"
)

type AdminController struct {
	beego.Controller
}

// @Title RegisterAdmin
// @Description register admin
// @Param	token	header true	string	"token"
// @Param	admin	body 	requestbody.AdminRequest	true	"Admin"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /create [post]
func (this *AdminController) RegisterAdmin() {
	defer this.ServeJSON()
	var err error
	body := requestbody.AdminRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = adminservice.RegisterAdminAccount(body)
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


// @Title update
// @Description register admin
// @Param	token	header true	string	"token"
// @Param	admin	body 	requestbody.AccountPutRequest	true	"Admin"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /update [put]
func (this *AdminController) UpdateAdmin() {
	defer this.ServeJSON()
	var err error
	body := requestbody.AccountPutRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = adminservice.UpdateAdminAccount(body)
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

// @Title RegisterAdmin
// @Description register admin
// @Param	token	header true	string	"token"
// @Param	admin	body 	requestbody.AccountLoginRequest	true	"Admin"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /login [post]
func (this *AdminController) LoginAdmin() {
	defer this.ServeJSON()
	var err error
	body := requestbody.AccountLoginRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	userId, err := adminservice.LoginAdmin(body)
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
	token, err := utils.CreateToken(userId, utils.ACCOUNT_ADMIN)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.ErrUnknown)
		return
	}
	//err = redisservice.CreateAuth(userId, token)
	//if err != nil {
	//	this.Data["json"] = responseservice.GetCommonErrorResponse(responses.ErrUnknown)
	//	return
	//}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(tokens)
}

// @Title GetAdmins
// @Description get roles
// @Param	token	header 	string true	"token"
// @Param	pos	query	int64	false	"Position"
// @Param	count	query	int64	false	"Count"
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *AdminController) Get() {
	defer this.ServeJSON()
	pos, err := this.GetInt32("pos", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	count, err := this.GetInt32("count",0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	if pos == 0 && count == 0 {
		this.GetAll()
	} else {
		this.GetPaginate(pos, count)
	}
}

func (this *AdminController) GetAll() {
	roles, err := adminservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(roles)
}

func (this *AdminController) GetPaginate(pos, count int32) {
	var err error
	roles, totalCount, err := adminservice.GetPaginate(pos, count)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseArray(roles, totalCount)
}


// @Title GetAdmins
// @Description get roles
// @Param	token	header	string true	"token"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /author-info [get]
func (this *AdminController) GetAuthorInfoForAdmin() {
	defer this.ServeJSON()
	adminId, err := this.GetInt32("adminId", 0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	authorInfo, err := adminservice.GetAuthorInfoForAdmin(adminId)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(authorInfo)
}



