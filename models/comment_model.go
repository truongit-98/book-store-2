package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Comment struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Content string `gorm:"type:text" json:"content"`
	CommentTime int64 `json:"comment_time"`
	Rate *uint `json:"rate"`
	BookID *uint ` json:"book_id"`
	CustomerID *uint `json:"customer_id"`
}

func (c Comment) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (c Comment) Exists() (exist bool, err error) {
	tx := db.First(&c, c.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (c Comment) GetBySort(sort string) (items interface{}, err error) {
	var comments []Comment
	tx := db.Order(sort).Find(&comments)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = comments
	return
}

func (c Comment) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var comments []Comment
	tx := db.Where(conds, params).Find(&comments)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = comments
	return}

func (c Comment) Create() (err error) {
	tx := db.Create(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (c Comment) GetById(id uint) (item interface{}, err error) {
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

func (c Comment) GetAll() (items interface{}, err error) {
	var comments []Comment
	tx := db.Find(&comments)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = comments
	return
}

func (c Comment) Update() (err error) {
	tx := db.Save(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (c Comment) Remove() (err error) {
	tx := db.Delete(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}