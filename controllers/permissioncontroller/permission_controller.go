package permissioncontroller

import (
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/permissionservice"
	"BookStore/services/responseservice"
	"encoding/json"
	"github.com/astaxie/beego"
)

type PermissionController struct {
	beego.Controller
}

// @Title CreatPermission
// @Description create new permission by token
// @Param	token		header 	string true	"token"
// @Param	permission	body 	requestbody.PermissionRequest	true	"Permission"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /create [post]
func (this *PermissionController) CreatePermission() {
	defer this.ServeJSON()
	var err error
	body := requestbody.PermissionRequest{}
	err = json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = permissionservice.CreateNewPermission(body)
	if err != nil {
		if _, ok := responses.MapDescription[err]; ok {
			this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		} else {
			this.Data["json"] = responses.ResponseCommonSingle{
				Data: map[string]interface{}{
					"message": err.Error(),
				},
				Error: responses.NewErr(responses.UnSuccess),
			}
		}
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponse()
}

// @Title GetPermissions
// @Description get permissions
// @Param	token	header 	string true	"token"
// @Param	pos	query	int64	false	"Position"
// @Param	count	query	int64	false	"Count"
// @Success 200 {object} responses.ResponseCommonArray
// @router / [get]
func (this *PermissionController) Get() {
	defer this.ServeJSON()
	pos, err := this.GetInt32("pos" ,0)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	count, err := this.GetInt32("count", 0)
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

func (this *PermissionController) GetAll() {
	roles, err := permissionservice.GetAll()
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(roles)

}

func (this *PermissionController) GetPaginate(pos, count int32) {
	var err error
	roles, totalCount, err := permissionservice.GetPaginate(pos, count)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseArray(roles, totalCount)
}

// @Title GetPermissionsForRole
// @Description get permissions for role
// @Param	token	header 	string true	"token"
// @Param	roleId	path	uint32	true	"role"
// @Success 200 {object} responses.ResponseCommonArray
// @router /role/:roleId [get]
func (this *PermissionController) GetPermissionsForRole() {
	defer this.ServeJSON()
	roleId, err := this.GetUint32(":roleId")
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(responses.BadRequest)
		return
	}
	data, err := permissionservice.GetPermissionsForRole(uint(roleId))
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponseArray(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponseWithData(data)
}
// @Title RemovePermission
// @Description remove permission
// @Param	token	header 	string true	"token"
// @Param	id	path 	uint32 true	"permissionID"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /delete/:id [delete]
func (this *PermissionController) RemovePermission() {
	defer this.ServeJSON()
	permissionId, err := this.GetInt64(":id")
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = permissionservice.RemovePermission(uint(permissionId))
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(err)
		return
	}
	this.Data["json"] = responseservice.GetCommonSucceedResponse()
}

// @Title UpdatePermission
// @Description update a permission
// @Param	token	header 	string true	"token"
// @Param	body	body	requestbody.PermissionPutRequest	true	"body"
// @Success 200 {object} responses.ResponseCommonSingle
// @router /update [put]
func (this *PermissionController) Put() {
	defer this.ServeJSON()
	body := requestbody.PermissionPutRequest{}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &body)
	if err != nil {
		this.Data["json"] = responseservice.GetCommonErrorResponse(responses.BadRequest)
		return
	}
	err = permissionservice.UpdatePermission(body)
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
