package voucherservice

import (
	"BookStore/models"
	"BookStore/restapi/responses"
	"github.com/prometheus/common/log"
)

func GetAll() ([]models.Voucher, error) {
	data, err := models.GetAll(&models.Voucher{})
	if err != nil {
		log.Info(err)
		return nil, responses.ErrSystem
	}
	return data.([]models.Voucher), nil
}
