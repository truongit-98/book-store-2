package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CategoryName string `gorm:"size:50" json:"category_name"`
	Topics []*Topic `json:"topics"`
}

func (c Category) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (c Category) Exists() (exist bool, err error) {
	tx := db.First(&c, c.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (c Category) GetBySort(sort string) (items interface{}, err error) {
	var categories []Category
	tx := db.Order(sort).Find(&categories)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = categories
	return
}

func (c Category) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var categories []Category
	tx := db.Where(conds, params).Find(&categories)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = categories
	return}

func (c Category) Create() (err error) {
	tx := db.Create(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (c Category) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&c, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = c
	return
}

func (c Category) GetAll() (items interface{}, err error) {
	var categories []Category
	tx := db.Preload("topics").Order("category_name").Find(&categories)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = categories
	return
}

func (c Category) GetAllWithCount() (items interface{}, err error) {
	var categories []responses.CategoryResponse
	tx := db.Raw("exec USP_GetAllCategoryWithCount").Scan(&categories)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = categories
	return
}


func (c Category) Update() (err error) {
	tx := db.Save(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (c Category) Remove() (err error) {
	tx := db.Delete(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}


