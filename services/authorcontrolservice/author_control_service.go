package authorcontrolservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"fmt"
	"github.com/prometheus/common/log"
	"strings"
)
var methods = map[string]string{
	"GET": "", "HEAD": "", "POST": "", "PUT": "", "DELETE": "", "CONNECT": "", "OPTIONS": "", "TRACE": "", "PATCH": "",
}
func CreateNewControl(body requestbody.AuthorControlRequest) error {
	if _, ok := methods[body.Action]; ok {
		name := strings.Join([]string{strings.ToUpper(body.Name[0:1]), body.Name[1:]}, "")
		control := models.AuthorControl{
			Name: name,
			Action: body.Action,
		}
		err := models.Create(control)
		if err != nil {
			log.Info(err.Error())
			return responses.ErrSystem
		}
		return nil
	}
	return fmt.Errorf("method %s not supported", body.Action)
}

func UpdateControl(body requestbody.AuthorControlPutRequest) error {
	if exist, err := models.Exists(models.AuthorControl{ID: body.Id}); err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	} else if exist {
		name := strings.Join([]string{strings.ToUpper(body.Name[0:1]), body.Name[1:]}, "")
		control := models.AuthorControl{
			ID: body.Id,
			Name: name,
			Action: body.Action,
		}
		err := models.Update(control)
		if err != nil {
			log.Info(err.Error())
			return responses.ErrSystem
		}
		return nil
	}
	return responses.NotExisted
}



func GetAll() ([]models.AuthorControl, error) {
	result := make([]models.AuthorControl, 0)
	data, err := models.GetAll(models.AuthorControl{})
	if err != nil {
		log.Info(err.Error())
		return result, responses.ErrSystem
	}
	result = data.([]models.AuthorControl)
	return result, nil
}

func GetPaginate(pos, count int32)([]models.AuthorControl, int32, error)  {
	result := make([]models.AuthorControl, 0)
	data, totalCount, err := models.GetPaginate(models.AuthorControl{}, pos, count)
	if err != nil {
		log.Info(err.Error())
		return result, totalCount,  responses.ErrSystem
	}
	result = data.([]models.AuthorControl)
	return result, totalCount, nil
}

func RemoveControl(id uint) error {
	control := models.AuthorControl{
		ID: id,
	}
	err := models.Remove(&control)
	if err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	}
	return nil
}