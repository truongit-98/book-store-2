package publisherservice

import (
	"BookStore/models"
	"BookStore/restapi/responses"
	"github.com/prometheus/common/log"
)

func GetAllWithBookCount() (interface{}, error){
	publisher := models.Publisher{}
	items, err := publisher.GetAllWithCount()
	if err != nil {
		log.Info(err.Error())
		return nil, responses.ErrSystem
	}
	return items, nil
}
