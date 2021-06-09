package adminservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/data"
	"BookStore/restapi/responses"
	"BookStore/services/permissionservice"
	"errors"
	"github.com/prometheus/common/log"
)

const DEFAULT_PASSWORD = "admin"

func GetAll() ([]models.Admin, error){
	admin := models.Admin{}
	items, err := models.GetAll(&admin)
	if err != nil {
		log.Info(err.Error())
		return make([]models.Admin, 0), err
	}
	return items.([]models.Admin), nil
}
func GetPaginate(pos, count int32)([]models.Admin, int32, error)  {
	result := make([]models.Admin, 0)
	data, totalCount, err := models.GetPaginate(&models.Admin{}, pos, count)
	if err != nil {
		log.Info(err.Error())
		return result, totalCount,  responses.ErrSystem
	}
	result = data.([]models.Admin)
	return result, totalCount, nil
}

func RegisterAdminAccount(body requestbody.AdminRequest) error {
	conds := "email = ?"
	data, err := models.GetWithCondition(&models.Admin{}, conds, body.Email)
	if err != nil {
		log.Info(err.Error())
		return err
	}
	if len(data.([]models.Admin)) <= 0 {
		admin := models.Admin{
			Email: body.Email,
			Password: DEFAULT_PASSWORD,
			FullName: body.FullName,
			Phone: body.Phone,
		}
		err := models.Create(&admin)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return errors.New("email is already taken")
}


func LoginAdmin(body requestbody.AccountLoginRequest) (int, error) {
	uid, err := (&models.Admin{}).Login(body.Email, body.Password)
	if err != nil {
		log.Info(err.Error())
		return 0, responses.ErrSystem
	}
	if uid <= 0 {
		return -1, responses.NotExisted
	}
	return uid, nil
}



func UpdateAdminAccount(body requestbody.AccountPutRequest) error {
	admin := models.Admin{ID: body.Id}
	if exist, err := models.Exists(&admin); err != nil {
		log.Info(err.Error())
		return err
	} else if exist {
		newAdmin := models.Admin{
			FullName: body.FullName,
			Address: body.Address,
			Phone: body.Phone,
		}
		err = models.Update(&newAdmin)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return responses.NotExisted
}

func GetAdminInfo(adminId int32) (*models.Admin, error) {
	adminInfo, err := (&models.Admin{}).GetAdminInfo(uint(adminId))
	if err != nil {
		log.Info(err)
		return nil, responses.ErrSystem
	}
	return adminInfo.(*models.Admin), nil
}

func GetAuthorInfoForAdmin(adminId int32) (*data.AdminResponse, error) {
	var resp = &data.AdminResponse{}
	info, err := GetAdminInfo(adminId)
	if err != nil {
		return nil, err
	}
	resp.ID = info.ID
	resp.Email = info.Email
	resp.FullName = info.FullName
	resp.Phone = info.Phone
	var respRoles = make([]*data.Role, 0)
	for _, r := range info.AdminRoles {
		var role = &data.Role{
			RoleID: int(r.RoleID),
		}
		permissions, err := permissionservice.GetPermissionsForRole(r.RoleID)
		if err != nil {
			return nil, err
		}
		for _, p := range permissions {
			var perms = &data.Permission{}
			perms.PermissionID = int(p.ID)
			perms.Name = p.Name
			perms.Path = p.Path
			for _, rp := range p.RolePerms{
				for _, rpc := range rp.RolePermControls{
					item, err := (models.AuthorControl{}).GetById(rpc.AuthorControlID)
					if err != nil {
						return nil, err
					}
					control := item.(models.AuthorControl)
					perms.Actions = append(perms.Actions, &data.Action{ActionID: int(control.ID), Action: control.Action})
				}
			}
			role.Permissions = append(role.Permissions, perms)
		}
		respRoles = append(respRoles, role)
	}
	resp.Roles = respRoles
	return resp, nil
}