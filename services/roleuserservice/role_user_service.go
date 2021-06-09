package roleuserservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"

	"github.com/prometheus/common/log"
)

func AddMultiRoleForUser(body requestbody.MultiRoleUserRequest) error {
	if exist, err := models.Exists(&models.Admin{ID: body.UserId}); err != nil {
		log.Info(err.Error())
		return responses.ErrSystem 
	} else if exist {
		for _, r := range body.Roles {
			if exist, err := models.Exists(&models.Role{ID: r.RoleId}); err != nil {
				log.Info(err.Error())
				return responses.ErrSystem 
			} else if exist {
				data, err := models.GetWithCondition(&models.AdminRole{}, &models.AdminRole{AdminID: body.UserId, RoleID: r.RoleId}) 
				if err != nil {
					log.Info(err.Error())
					return responses.ErrSystem 
				}
				roleUser := data.([]models.AdminRole)
				if len(roleUser) > 0 {
					if !r.Active {
						err = models.Remove(&roleUser[0])
						if err != nil {
							log.Info(err.Error())
							return responses.ErrSystem
						}
					}
				} else if r.Active == true {
					err := models.Create(&models.AdminRole{AdminID: body.UserId, RoleID: r.RoleId})
					if err != nil {
						log.Info(err.Error())
						return responses.ErrSystem 
					}
				} 
				continue
			} 
			
		}
		return nil
	}
	return responses.NotExisted
}


func GetAllRoleForUser(userId uint) ([]models.AdminRole, error) {
	conds := models.AdminRole{AdminID: userId}
	data, err := models.GetWithCondition(&models.AdminRole{}, conds)
	if err != nil {
		log.Info(err.Error())
		return make([]models.AdminRole, 0), responses.ErrSystem
	}
	return data.([]models.AdminRole), nil

}

func GetAll() ([]models.AdminRole, error) {
	result := make([]models.AdminRole, 0)
	role := models.AdminRole{}
	data, err := models.GetAll(&role)
	if err != nil {
		log.Info(err.Error())
		return result, responses.ErrSystem
	}
	result = data.([]models.AdminRole)
	return result, nil
}
