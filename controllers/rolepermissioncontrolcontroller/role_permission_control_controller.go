package rolepermissioncontrolcontroller

import (
	"BookStore/services/responseservice"
	"BookStore/services/rolepermissioncontrolservice"
	"github.com/astaxie/beego"
)

type RolePermsControlController struct {
	beego.Controller
}

// @Title GetAllRolePermControls
// @Description Get All RolePermControls
// @Param	token	header 	string true	"token"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /all [get]
func (this *RolePermsControlController) GetAllRolePermControls() {
	defer this.ServeJSON()
	data, err := rolepermissioncontrolservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)
}