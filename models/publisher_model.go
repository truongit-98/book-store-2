package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type Publisher struct {
	ID uint                      `gorm:"primaryKey;autoIncrement" json:"id"`
	PublisherName string `gorm:"size:50" json:"publisher_name"`
	Address string       `gorm:"size:256" json:"address"`
	Country string       `gorm:"size:50" json:"country"`
	Email string         `gorm:"size:50;not null" json:"email"`
	Phone string       `gorm:"size:10" json:"phone"`
	Specifications []*BookDetail ` json:"specifications"`
}

func (p Publisher) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (p Publisher) Exists() (exist bool, err error) {
	tx := db.First(&p, p.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (p Publisher) GetBySort(sort string) (items interface{}, err error) {
	var publishers []Publisher
	tx := db.Order(sort).Find(&publishers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = publishers
	return
}

func (p Publisher) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var publishers []Publisher
	tx := db.Where(conds, params).Find(&publishers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = publishers
	return}

func (p Publisher) Create() (err error) {
	tx := db.Create(&p)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (p Publisher) Create2() (id uint, err error) {
	var count int 
	_ = db.Raw("select count(id) from publishers").Scan(&count)
	if count < 20 {
		tx:= db.Where("publisher_name = ?", p.PublisherName).First(&p)
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
		tx := db.Create(&p)
		if tx.Error != nil {
			err = responses.ErrSystem
			return
		}
		return p.ID, nil
	}
	return p.ID, nil
	} else {
		rand.Seed(time.Now().UnixNano())
		min := 1
		max := 20
		id := rand.Intn(max-min+1) + min
		return uint(id), nil
	}
	
}

func (p Publisher) GetById(id uint) (item interface{}, err error) {
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

func (p Publisher) GetAll() (items interface{}, err error) {
	var publishers []Publisher
	tx := db.Order("publisher_name").Find(&publishers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = publishers
	return
}

func (p Publisher) GetAllWithCount() (items interface{}, err error) {
	var publishers []responses.PublisherResponseWithCountBook
	tx := db.Raw("exec USP_GetAllPublisherWithCount").Scan(&publishers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = publishers
	return
}

func (p Publisher) Update() (err error) {
	tx := db.Save(&p)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (p Publisher) Remove() (err error) {
	tx := db.Delete(&p)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

