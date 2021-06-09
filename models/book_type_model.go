package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type BookType struct {
	ID uint                      `json:"id" xml:"id"`
	Type string          `gorm:"size:50" json:"type"`
	TypeName string      `gorm:"size:50" json:"type_name" xml:"type_name"`
	Episodes int       `json:"episodes"`
	Books []*Book                `json:"books"`
	ProductDetails []*BookDetail `json:"product_details"`
	Awards []*Award `gorm:"many2many:book_type_awards;joinForeignKey:BookTypeID" json:"awards"`
	Authors []*Author `gorm:"many2many:book_type_authors;joinForeignKey:BookTypeID" json:"authors"`
}

func (b *BookType) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (b *BookType) Exists() (exist bool, err error) {
	tx := db.First(&b, b.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (b Book) BookType(sort string) (items interface{}, err error) {
	var books []BookType
	tx := db.Order(sort).Find(&books)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = books
	return
}

func (b *BookType) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var books []BookType
	tx := db.Where(conds, params).Find(&books)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = books
	return
}

func (b *BookType) Create() (err error) {
	tx := db.Create(&b)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (b *BookType) CreateV2() (uint, error) {
	tx := db.Create(&b)
	if tx.Error != nil {
		return 0, tx.Error
	}
	return b.ID, nil
}

func (b *BookType) Create2() (id uint, err error) {
	tx := db.Where("type_name = ?", b.TypeName).First(&b)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound){
		tx := db.Create(&b)
		if tx.Error != nil {
			err = responses.ErrSystem
			return
		}
		return b.ID, nil
	}
	return b.ID, nil
}

func (b *BookType) GetById(id uint) (item interface{}, err error) {
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
func (b *BookType) GetBySort(sort string) (items interface{}, err error) {
	var books []BookType
	tx := db.Order(sort).Find(&books)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = books
	return
}

func (b *BookType) GetAll() (items interface{}, err error) {
	var books []BookType
	tx := db.Preload("Comments").Find(&books)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = books
	return
}


func (b *BookType) Update() (err error) {
	tx := db.Save(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (b *BookType) Remove() (err error) {
	tx := db.Delete(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}


