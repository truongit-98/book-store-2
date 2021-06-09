package UserController

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/customerservice"
	"BookStore/services/responseservice"
	"BookStore/util"
	"encoding/json"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

// @Title RegisterUser
// @Description register user
// @Param	user	body 	requestbody.AccountRequest	true	"user"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /register [post]
func (this *UserController) RegisterUser() {
	defer this.ServeJSON()
	var err error
	body := requestbody.AccountRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = customerservice.RegisterUserAccount(body)
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

// @Title LoginUser
// @Description login user
// @Param	admin	body 	requestbody.AccountLoginRequest	true	"Admin"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /login [post]
func (this *UserController) LoginUser() {
	defer this.ServeJSON()
	var err error
	body := requestbody.AccountLoginRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	userId, err := customerservice.LoginUser(body)
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
	token, err := utils.CreateToken(userId, utils.ACCOUNT_USER)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.ErrUnknown)
		return
	}
	// err = redisservice.CreateAuth(userId, token)
	// if err != nil {
	// 	this.Data["json"] = responseservice.GetCommonErrorResponse(responses.ErrUnknown)
	// 	return
	// }
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(tokens)
}

// @Title GetUsers
// @Description get users
// @Param	token	header 	string true	"token"
// @Param	pos	query	int64	false	"Position"
// @Param	count	query	int64	false	"Count"
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *UserController) Get() {
	defer this.ServeJSON()
	pos, err := this.GetInt32("pos")
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	count, err := this.GetInt32("count")
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

func (this *UserController) GetAll() {
	roles, err := customerservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(roles)
}

func (this *UserController) GetPaginate(pos, count int32) {
	var err error
	roles, totalCount, err := customerservice.GetPaginate(pos, count)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseArray(roles, totalCount)
}