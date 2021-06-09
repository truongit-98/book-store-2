package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Topic struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	TopicName string `gorm:"size:50" json:"topic_name"`
	CategoryID *uint `json:"topic_id"`
	Specifications []*BookDetail `json:"specifications"`
}

func (t Topic) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (t Topic) Exists() (exist bool, err error) {
	tx := db.First(&t, t.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (t Topic) GetBySort(sort string) (items interface{}, err error) {
	var topics []Topic
	tx := db.Order(sort).Find(&topics)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = topics
	return
}

func (t Topic) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var topics []Topic
	tx := db.Where(conds, params).Find(&topics)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = topics
	return}

func (t Topic) Create() (err error) {
	tx := db.Create(&t)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (t Topic) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&t, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = t
	return
}

func (t Topic) GetAll() (items interface{}, err error) {
	var topics []Topic
	tx := db.Find(&topics)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = topics
	return
}

func (t Topic) Update() (err error) {
	tx := db.Save(&t)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (t Topic) Remove() (err error) {
	tx := db.Delete(&t)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

