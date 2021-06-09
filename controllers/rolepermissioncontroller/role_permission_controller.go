package rolepermissioncontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/responseservice"
	"BookStore/services/rolepermissionservice"
	"encoding/json"
	"github.com/astaxie/beego"
)

type RolePermissionController struct {
	beego.Controller
}

// @Title GetRolePermissions
// @Description get role_permissions
// @Param	token	header 	string true	"token"
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *RolePermissionController) GetRolePermissions() {
	defer this.ServeJSON()
	roles, err := rolepermissionservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(roles)
}

// @Title AddMultiPermissionToRole
// @Description add permission for role
// @Param	token	header 	string true	"token"
// @Param	body	body	requestbody.MultiRolePermissionRequest	true	"body"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /create/multi [post]
func (this *RolePermissionController) AddMultiPermsToRole() {
	defer this.ServeJSON()
	body := requestbody.MultiRolePermissionRequest{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = rolepermissionservice.AddMultiPermissionToRole(body)
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
