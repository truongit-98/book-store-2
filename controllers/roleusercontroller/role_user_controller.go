package roleusercontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/responseservice"
	"BookStore/services/roleuserservice"
	"encoding/json"
	"github.com/astaxie/beego"
)

type RoleUserController struct {
	beego.Controller
}

// @Title GetRoleUsers
// @Description get role_permissions
// @Param	token	header 	string true	"token"
// @Success 200 {object} responses.ResponseCommonArray
// @router /all [get]
func (this *RoleUserController) GetRoleUsers() {
	defer this.ServeJSON()
	roles, err := roleuserservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(roles)
}

// @Title GetRoleUsers
// @Description get role_permissions
// @Param	token	header 	string true	"token"
// @Param	userId path	int32	true "UserID"
// @Success 200 {object} responses.ResponseCommonArray
// @router /user/:userId [get]
func (this *RoleUserController) GetRoleUsersForUser() {
	defer this.ServeJSON()
	userId, err := this.GetInt32(":userId")
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	data, err := roleuserservice.GetAllRoleForUser(uint(userId))
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)
}

// @Title AddRoleForUser
// @Description Add role for user
// @Param	token	header 	string true	"token"
// @Param	body	body	requestbody.MultiRoleUserRequest	true	"body"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /multi [post]
func (this *RoleUserController) AddMultiRoleForUser() {
	defer this.ServeJSON()
	body := requestbody.MultiRoleUserRequest{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = roleuserservice.AddMultiRoleForUser(body)
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