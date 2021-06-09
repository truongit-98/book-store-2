package rolepermissioncontrolservice

import (
	"BookStore/models"
	"BookStore/restapi/responses"
	"github.com/prometheus/common/log"
)


func GetAll() ([]models.RolePermissionControl, error) {
	result := make([]models.RolePermissionControl, 0)
	data, err := models.GetAll(&models.RolePermissionControl{})
	if err != nil {
		log.Info(err.Error())
		return result, responses.ErrSystem
	}
	result = data.([]models.RolePermissionControl)
	return result, nil
}

func GetRolePermsControlForRolePerms(rolePermsId uint) ([]models.RolePermissionControl, error) {
	result := make([]models.RolePermissionControl, 0)
	conds := "role_permissions_id = ?"
	data, err := models.GetWithCondition(&models.RolePermissionControl{}, conds, rolePermsId)
	if err != nil {
		log.Info(err.Error())
		return make([]models.RolePermissionControl, 0), responses.ErrSystem
	}
	result = data.([]models.RolePermissionControl)
	return result, nil
}


func GetRolePermsControlByRolePermsAndControl(rolePermsId, controlId uint) (*models.RolePermissionControl, error) {
	conds := models.RolePermissionControl{RolePermissionsID: rolePermsId, AuthorControlID: controlId}
	data, err := models.GetWithCondition(&models.RolePermissionControl{}, conds)
	if err != nil {
		log.Info(err.Error())
		return nil, responses.ErrSystem
	}
	result := data.([]models.RolePermissionControl)

	if len(result) > 0 {
		return &result[0], nil
	}
	return nil, nil
}


func CreateRolePermsControl(rolePermsControl models.RolePermissionControl) error {
	err := models.Create(&rolePermsControl)
	if err != nil {
		log.Info(err)
		return responses.ErrSystem
	}
	return nil
}