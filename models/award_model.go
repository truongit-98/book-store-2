package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Award struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Award string `gorm:"size:255" json:"award"`
	Types []*BookType `gorm:"many2many:book_type_awards;joinForeignKey:AwardID" json:"types"`
}

func (a Award) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (a Award) Exists() (exist bool, err error) {
	tx := db.First(&a, a.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (a Award) GetBySort(sort string) (items interface{}, err error) {
	var awards []Award
	tx := db.Order(sort).Find(&awards)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = awards
	return
}

func (a Award) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var awards []Award
	tx := db.Where(conds, params).Find(&awards)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = awards
	return
}

func (a Award) Create() (err error) {
	tx := db.Create(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (a Award) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&a, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = a
	return
}

func (a Award) GetAll() (items interface{}, err error) {
	var awards []Award
	tx := db.Find(&awards)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = awards
	return
}

func (a Award) Update() (err error) {
	tx := db.Save(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (a Award) Remove() (err error) {
	tx := db.Delete(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

