package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type OrderDetail struct {
	ID uint `json:"id"`
	OrderID uint `json:"order_id"`
	BookID uint `json:"book_id"`
	Quantity int `json:"quantity"`
	Price float64 `json:"price"`
}

func (o* OrderDetail) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (o* OrderDetail) Exists() (exist bool, err error) {
	tx := db.Where("book_id = ? AND order_id = ?", o.BookID, o.OrderID).First(&o)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (o* OrderDetail) GetBySort(sort string) (items interface{}, err error) {
	var details []OrderDetail
	tx := db.Order(sort).Find(&details)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = details
	return
}

func (o* OrderDetail) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var details []OrderDetail
	tx := db.Where(conds, params).Find(&details)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = details
	return}

func (o* OrderDetail) Create() (err error) {
	tx := db.Create(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (o* OrderDetail) GetById(id uint) (item interface{}, err error) {
	return nil, errors.New("not supported method")
}

func (o* OrderDetail) GetAll() (items interface{}, err error) {
	var details []OrderDetail
	tx := db.Find(&details)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = details
	return
}

func (o* OrderDetail) Update() (err error) {
	tx := db.Save(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (o* OrderDetail) Remove() (err error) {
	tx := db.Delete(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}
