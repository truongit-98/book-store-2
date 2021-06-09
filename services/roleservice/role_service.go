package roleservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"BookStore/services/roleuserservice"
	"errors"
	"gorm.io/gorm"

	"github.com/prometheus/common/log"
)

func CreateNewRole(body requestbody.RoleRequest) (uint, error) {
	role := models.Role{}
	role.Name = body.Name
	err := models.Create(&role)
	if err != nil {
		log.Info(err.Error())
		return 0, responses.ErrSystem
	}
	return role.ID, nil
}

func UpdateRole(body requestbody.RolePutRequest) error {
	item, err := models.GetById(&models.Role{}, body.RoleId)
	if err != nil {
		log.Info(err.Error())
		if errors.Is(err, gorm.ErrRecordNotFound){
			return responses.NotExisted
		}
		return responses.ErrSystem
	}
	oldRole := item.(*models.Role)
	oldRole.Name = body.Name
	err = models.Update(oldRole)
	if err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	}
	return nil
}

func RemoveRole(id uint) error {
	role := models.Role{
		ID: id,
	}
	err := models.Remove(&role)
	if err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	}
	return nil
}

func GetAll() ([]models.Role, error) {
	result := make([]models.Role, 0)
	role := models.Role{}
	data, err := models.GetAll(&role)
	if err != nil {
		log.Info(err.Error())
		return result, responses.ErrSystem
	}
	result = data.([]models.Role)
	return result, nil
}

func GetPaginate(pos, count int32)([]models.Role, int32, error)  {
	result := make([]models.Role, 0)
	role := models.Role{}
	data, totalCount, err := models.GetPaginate(&role, pos, count)
	if err != nil {
		log.Info(err.Error())
		return result, totalCount,  responses.ErrSystem
	}
	result = data.([]models.Role)
	return result, totalCount, nil
}

func GetRoleInfo(roleId uint) (*models.Role, error){
	data, err := models.GetById(&models.Role{}, roleId)
	if err != nil {
		log.Info(err.Error())
		return nil,responses.ErrSystem
	}
	//log.Info(data,"[GetRoleInfo]")
	role := data.(*models.Role)
	return role, nil

}

func GetRolesForUser(userId uint) ([]*models.Role, error) {
	result := make([]*models.Role, 0)
	userRoles, err := roleuserservice.GetAllRoleForUser(userId)
	if err != nil {
		return make([]*models.Role, 0), err
	}
	for _, e := range userRoles {
		role, err := models.GetById(&models.Role{}, e.RoleID)
		if err != nil {
			log.Info(err.Error())
			return make([]*models.Role, 0), err
		}
		result = append(result, role.(*models.Role))
	}
	return result, nil
}

