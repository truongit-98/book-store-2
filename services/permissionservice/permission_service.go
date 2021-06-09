package permissionservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/rolepermissionservice"
	"github.com/prometheus/common/log"
)

func CreateNewPermission(body requestbody.PermissionRequest) error {
	perm := models.Permission{
		Name: body.Name,
		Path: body.Path,
	}
	err := models.Create(perm)
	if err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	}
	return nil

}

func UpdatePermission(body requestbody.PermissionPutRequest) error {
	if exist, err := models.Exists(models.Permission{ID: body.PermissionId}); err != nil {
		return responses.ErrSystem
	} else if exist {
		perm := models.Permission{
			ID:     body.PermissionId,
			Name:   body.Name,
			Path: body.Path,
		}
		err := models.Update(perm)
		if err != nil {
			log.Info(err.Error())
			return responses.ErrSystem
		}
		return nil
	}
	return responses.NotExisted
}

func RemovePermission(id uint) error {
	perm := models.Permission{
		ID: id,
	}
	err := models.Remove(perm)
	if err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	}
	return nil
}

func GetAll() ([]models.Permission, error) {
	result := make([]models.Permission, 0)
	perm := models.Permission{}
	data, err := models.GetAll(perm)
	if err != nil {
		log.Info(err.Error())
		return result, responses.ErrSystem
	}
	result = data.([]models.Permission)
	return result, nil
}

func GetPaginate(pos, count int32)([]models.Permission, int32, error)  {
	result := make([]models.Permission, 0)
	perm := models.Permission{}
	data, totalCount, err := models.GetPaginate(perm, pos, count)
	if err != nil {
		log.Info(err.Error())
		return result, totalCount,  responses.ErrSystem
	}
	result = data.([]models.Permission)
	return result, totalCount, nil
}

func GetPermissionsForRole(roleId uint) ([]models.Permission, error) {
	result := make([]models.Permission, 0)
	routers, err := rolepermissionservice.GetRolePermsForRole(roleId)
	if err != nil {
		return make([]models.Permission, 0), err
	}
	for _, e := range routers {
		perm, err := models.GetById(models.Permission{}, e.PermissionID)
		if err != nil {
			log.Info(err.Error())
			return make([]models.Permission, 0), err
		}
		result = append(result, perm.(models.Permission))
	}
	return result, nil
}
