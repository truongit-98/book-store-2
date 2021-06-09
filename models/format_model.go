package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Format struct {
	ID uint                      `gorm:"primaryKey;autoIncrement" json:"id"`
	TypeFormat string    `gorm:"size:50" json:"type_format"`
	Specifications []*BookDetail `json:"specifications"`
}

func (f Format) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (f Format) Exists() (exist bool, err error) {
	tx := db.First(&f, f.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (f Format) GetBySort(sort string) (items interface{}, err error) {
	var formats []Format
	tx := db.Order(sort).Find(&formats)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = formats
	return
}

func (f Format) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var formats []Format
	tx := db.Where(conds, params).Find(&formats)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = formats
	return}

func (f Format) Create() (err error) {
	tx := db.Create(&f)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (f Format) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&f, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = f
	return
}

func (f Format) GetAll() (items interface{}, err error) {
	var formats []Format
	tx := db.Find(&formats)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = formats
	return
}

func (f Format) Update() (err error) {
	tx := db.Save(&f)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (f Format) Remove() (err error) {
	tx := db.Delete(&f)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}
