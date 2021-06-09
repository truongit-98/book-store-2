package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Payment struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Method string `gorm:"size:50" json:"method"`
	Orders []*Order `json:"orders"`
}

func (p Payment) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (p Payment) Exists() (exist bool, err error) {
	tx := db.First(&p, p.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (p Payment) GetBySort(sort string) (items interface{}, err error) {
	var payments []Payment
	tx := db.Order(sort).Find(&payments)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = payments
	return
}

func (p Payment) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var payments []Payment
	tx := db.Where(conds, params).Find(&payments)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = payments
	return}

func (p Payment) Create() (err error) {
	tx := db.Create(&p)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (p Payment) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&p, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = p
	return
}

func (p Payment) GetAll() (items interface{}, err error) {
	var payments []Payment
	tx := db.Find(&payments)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = payments
	return
}

func (p Payment) Update() (err error) {
	tx := db.Save(&p)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (p Payment) Remove() (err error) {
	tx := db.Delete(&p)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

