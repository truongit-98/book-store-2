package rolepermissionservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/rolepermissioncontrolservice"
	"errors"
	"github.com/prometheus/common/log"
	"gorm.io/gorm"
)

func AddMultiPermissionToRole(body requestbody.MultiRolePermissionRequest) error {
	if exist, err := models.Exists(&models.Role{ID: body.RoleId}); err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	} else if exist {
		err = models.Transaction(func(tx *gorm.DB) error {
			for _, p := range body.Permissions{
				if exist, err := models.Exists(models.Permission{ID: p.PermissionId}); err != nil {
					log.Info(err.Error())
					return responses.ErrSystem
				} else if !exist {
					return responses.BadRequest
				}
				var rolePermsId = uint(0)
				// check role-perms exist
				rolePerms, err := GetRolePermsForRoleAndPermission(body.RoleId, p.PermissionId)
				if err != nil {
					return responses.ErrSystem
				}
				if rolePerms == nil {
					rolePerms = &models.RolePermissions{RoleID: body.RoleId, PermissionID: p.PermissionId}
					rolePermsId, err  = CreateRolePerms(rolePerms)
					if err != nil {
						return err
					}
				} else {
					rolePermsId = rolePerms.ID
				}
				for _, c := range p.Controls {
					if exist, err := models.Exists(models.AuthorControl{ID: c.ControlId}); err != nil {
						log.Info(err)
						return responses.ErrSystem
					} else if !exist {
						return responses.BadRequest
					}
					rolePermsControl, err := rolepermissioncontrolservice.GetRolePermsControlByRolePermsAndControl(rolePermsId, c.ControlId)
					if err != nil {
						return err
					}
					if rolePermsControl == nil {
						if c.Active {
							err = models.Create(&models.RolePermissionControl{RolePermissionsID: rolePermsId, AuthorControlID: c.ControlId})
							if err != nil {
								log.Info(err)
								return responses.ErrSystem
							}
							if c.ControlId == 2 || c.ControlId == 3 || c.ControlId == 4 {
								controlGet, err := models.GetById(&models.RolePermissionControl{}, 1)
								if err != nil && !errors.Is(err, gorm.ErrRecordNotFound){
									log.Info(err)
									return responses.ErrSystem
								}
								if controlGet == nil {
									err = models.Create(&models.RolePermissionControl{RolePermissionsID: rolePermsId, AuthorControlID: 1})
									if err != nil {
										log.Info(err)
										return responses.ErrSystem
									}
								}
							}
						}
					} else if !c.Active{
						err = models.Remove(rolePermsControl)
						if err != nil {
							log.Info(err)
							return responses.ErrSystem
						}
						if c.ControlId == 1 {
							conds := "role_permissions_id = ?"
							rolepermcontrols, err := models.GetWithCondition(&models.RolePermissionControl{},conds, rolePermsId)
							if err != nil {
								log.Info(err)
								return responses.ErrSystem
							}
							for _, rpc := range rolepermcontrols.([]models.RolePermissionControl){
								err = models.Remove(&rpc)
								if err != nil {
									log.Info(err)
									return responses.ErrSystem
								}
							}
						}
					}
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
		return nil
	}
	return responses.BadRequest
}


func GetRolePermsForRole(roleId uint) ([]models.RolePermissions, error) {
	conds := "role_id = ?"
	data, err := models.GetWithCondition(&models.RolePermissions{}, conds, roleId)
	if err != nil {
		log.Info(err.Error())
		return make([]models.RolePermissions, 0), responses.ErrSystem
	}
	return data.([]models.RolePermissions), nil
}

func GetRolePermsForRoleAndPermission(roleId, permissionId uint) (*models.RolePermissions, error) {
	conds := models.RolePermissions{RoleID: roleId, PermissionID: permissionId}
	data, err := models.GetWithCondition(&models.RolePermissions{}, conds)
	if err != nil {
		log.Info(err.Error())
		return nil, responses.ErrSystem
	}
	if len(data.([]models.RolePermissions)) > 0 {
		return &data.([]models.RolePermissions)[0], nil
	}
	return nil, nil
}


func GetAll() ([]models.RolePermissions, error) {
	result := make([]models.RolePermissions, 0)
	data, err := models.GetAll(&models.RolePermissions{})
	if err != nil {
		log.Info(err.Error())
		return result, responses.ErrSystem
	}
	result = data.([]models.RolePermissions)
	return result, nil
}

func GetPaginate(pos, count int32)([]models.RolePermissions, int32, error)  {
	result := make([]models.RolePermissions, 0)
	data, totalCount, err := models.GetPaginate(&models.RolePermissions{}, pos, count)
	if err != nil {
		log.Info(err.Error())
		return result, totalCount,  responses.ErrSystem
	}
	result = data.([]models.RolePermissions)
	return result, totalCount, nil
}

func CreateRolePerms(rolePersm *models.RolePermissions) (uint, error) {
	err := models.Create(rolePersm)
	if err != nil {
		return 0, err
	}
	return rolePersm.ID, nil
}
