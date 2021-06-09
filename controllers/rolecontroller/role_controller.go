package rolecontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/responseservice"
	"BookStore/services/roleservice"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/prometheus/common/log"
)

type RoleController struct {
	beego.Controller
}

func (this *RoleController) setLocation(id uint) {
	this.Ctx.ResponseWriter.Header().Add("location", fmt.Sprintf("http://localhost:8088/v1/api/role/detail/%d", id))
}

// @Title CreateRole
// @Description create new role by token
// @Param	token	header true	string	"token"
// @Param	role	body 	requestbody.RoleRequest	true	"Role"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /create [post]
func (this *RoleController) CreateRole() {
	defer this.ServeJSON()
	var err error
	body := requestbody.RoleRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	id, err := roleservice.CreateNewRole(body)
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
	this.setLocation(id)
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(map[string]interface{}{
		"id": id,
	})
}

// @Title GetRoles
// @Description get roles
// @Param	token	header 	string true	"token"
// @Param	pos	query	int64	false	"Position"
// @Param	count	query	int64	false	"Count"
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *RoleController) Get() {
	defer this.ServeJSON()
	pos, err := this.GetInt32("pos", 0)
	if err != nil {
		log.Info(err)
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	count, err := this.GetInt32("count", 0)
	if err != nil {
		log.Info(err)

		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	if pos == 0 && count == 0 {
		this.GetAll()
	} else {
		this.GetPaginate(pos, count)
	}
}

func (this *RoleController) GetAll() {
	roles, err := roleservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(roles)
}

func (this *RoleController) GetPaginate(pos, count int32) {
	var err error
	roles, totalCount, err := roleservice.GetPaginate(pos, count)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseArray(roles, totalCount)
}

// @Title GetRolesById
// @Description get roles by  id
// @Param	token	header 	string true	"token"
// @Param	roleId 	path	int32	true "roleId"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /detail/:roleId [get]
func (this *RoleController) GetRolesById() {
	defer this.ServeJSON()
	roleId, err := this.GetInt32(":roleId")
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	data, err := roleservice.GetRoleInfo(uint(roleId))
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

// @Title GetRolesForUser
// @Description get roles for user
// @Param	token	header 	string true	"token"
// @Param	userId 	path	int32	true "UserID"
// @Success 200 {object} responses.ResponseCommonArray
// @router /user/:userId [get]
func (this *RoleController) GetRolesForUser() {
	defer this.ServeJSON()
	userId, err := this.GetInt32(":userId")
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	data, err := roleservice.GetRolesForUser(uint(userId))
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)

}

// @Title UpdateRole
// @Description update a role
// @Param	token	header 	string true	"token"
// @Param	body	body	requestbody.RolePutRequest	true	"body"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /update [put]
func (this *RoleController) Put() {
	defer this.ServeJSON()
	body := requestbody.RolePutRequest{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = roleservice.UpdateRole(body)
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

// @Title Removerole
// @Description remove role by  id
// @Param	token	header 	string true	"token"
// @Param	roleId 	path	int32	true "roleId"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /delete/:roleId [delete]
func (this *RoleController) Removerole() {
	defer this.ServeJSON()
	roleId, err := this.GetInt32(":roleId")
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = roleservice.RemoveRole(uint(roleId))
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponse()

}