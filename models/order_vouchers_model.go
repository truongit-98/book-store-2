package models

import (
	"BookStore/restapi/responses"
	"errors"
)

type OrderVouchers struct {
	ID uint `json:"id"`
	OrderID uint `json:"order_id" xml:"order_id"`
	VoucherID uint `json:"voucher_id" xml:"voucher_id"`
	Value float64 `json:"value xml:"value""`
}

func (o* OrderVouchers) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (o* OrderVouchers) Exists() (exist bool, err error) {
	//tx := db.Where("voucher_id = ? AND order_id = ?", o.VoucherID, o.OrderID).First(&o)
	//if err := tx.Error;  err != nil {
	//	if !errors.Is(err, gorm.ErrRecordNotFound) {
	//		return false, err
	//	}
	//	return false, nil
	//}
	return true, nil
}

func (o* OrderVouchers) GetBySort(sort string) (items interface{}, err error) {
	var details []OrderDetail
	tx := db.Order(sort).Find(&details)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = details
	return
}

func (o* OrderVouchers) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var details []OrderDetail
	tx := db.Where(conds, params).Find(&details)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = details
	return}

func (o* OrderVouchers) Create() (err error) {
	tx := db.Create(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (o* OrderVouchers) GetById(id uint) (item interface{}, err error) {
	return nil, errors.New("not supported method")
}

func (o* OrderVouchers) GetAll() (items interface{}, err error) {
	var details []OrderDetail
	tx := db.Find(&details)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = details
	return
}

func (o* OrderVouchers) Update() (err error) {
	tx := db.Save(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (o* OrderVouchers) Remove() (err error) {
	tx := db.Delete(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}