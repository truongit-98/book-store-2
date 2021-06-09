package paymentmethodservice

import (
	"BookStore/models"
	"BookStore/restapi/responses"
	"github.com/prometheus/common/log"
)

func GetAll() ([]models.Payment, error){
	payments, err := models.GetAll(models.Payment{})
	if err != nil {
		log.Info(err)
		return nil, responses.ErrSystem
	}
	return payments.([]models.Payment), nil
}

