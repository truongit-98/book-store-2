package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Delivery struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Method string `gorm:"size:50" json:"method"`
	Orders []*Order `gorm:"foreignKey:Delivery" json:"orders"`
}

func (d Delivery) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (d Delivery) Exists() (exist bool, err error) {
	tx := db.First(&d, d.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (d Delivery) GetBySort(sort string) (items interface{}, err error) {
	var deliveries []Delivery
	tx := db.Order(sort).Find(&deliveries)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = deliveries
	return
}

func (d Delivery) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var deliveries []Delivery
	tx := db.Where(conds, params).Find(&deliveries)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = deliveries
	return}

func (d Delivery) Create() (err error) {
	tx := db.Create(&d)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (d Delivery) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&d, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = d
	return
}

func (d Delivery) GetAll() (items interface{}, err error) {
	var deliveries []Delivery
	tx := db.Find(&deliveries)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = deliveries
	return
}

func (d Delivery) Update() (err error) {
	tx := db.Save(&d)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (d Delivery) Remove() (err error) {
	tx := db.Delete(&d)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}