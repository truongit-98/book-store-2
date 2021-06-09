package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type BookDetail struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Price	float64	`json:"price"`
	PriceCover     float64 `json:"price_cover"`
	NumberOfPage   int    `json:"number_of_page"`
	Height         float64   `json:"height"`
	Weight         float64 `json:"weight"`
	Width         float64 `json:"width"`
	Description    string  `gorm:"type:text" json:"description"`
	Language       string  `gorm:"size:20" json:"language"`
	Active bool `json:"active" xml:"active"`
	BookTypeID         uint `json:"book_type_id"`
	PublisherID    uint    `json:"publisher_id"`
	FormatID       uint    `json:"format_id"`
	TopicID uint                 `json:"topic_id"`
}

func (b *BookDetail) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (b *BookDetail) Exists() (exist bool, err error) {
	tx := db.First(&b, b.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (b *BookDetail) GetBySort(sort string) (items interface{}, err error) {
	var bookSpecs []BookDetail
	tx := db.Order(sort).Find(&bookSpecs)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookSpecs
	return
}

func (b *BookDetail) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var bookSpecs []BookDetail
	tx := db.Where(conds, params).Find(&bookSpecs)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookSpecs
	return
}

func (b *BookDetail) Create() (err error) {
	tx := db.Create(&b)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (b *BookDetail) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&b, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = b
	return
}

func (b *BookDetail) GetAll() (items interface{}, err error) {
	var bookSpecs []BookDetail
	tx := db.Find(&bookSpecs)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookSpecs
	return
}

func (b *BookDetail) Update() (err error) {
	tx := db.Save(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (b *BookDetail) Remove() (err error) {
	tx := db.Delete(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

