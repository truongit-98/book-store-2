package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Author struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	AuthorName string `gorm:"size:50" json:"author_name"`
	BirthDate int64 `json:"birth_date"`
	Nationality string `gorm:"size:20" json:"nationality"`
	Email string `gorm:"size:50" json:"email"`
	Phone string `gorm:"size:10" json:"phone"`
	Types []*BookType `gorm:"many2many:book_type_authors;joinForeignKey:AuthorID" json:"types"`
}

func (a Author) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (a Author) Exists() (exist bool, err error) {
	tx := db.First(&a, a.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (a Author) GetBySort(sort string) (items interface{}, err error) {
	var authors []Author
	tx := db.Order(sort).Find(&authors)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = authors
	return
}

func (a Author) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var authors []Author
	tx := db.Where(conds, params).Find(&authors)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = authors
	return
}

func (a Author) Create() (err error) {
	tx := db.Create(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}
func (a Author) Create2() (id uint, err error) {
	tx := db.Where("author_name = ?", a.AuthorName).First(&a)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound){
		tx := db.Create(&a)
		if tx.Error != nil {
			err = responses.ErrSystem
			return
		}
		return a.ID, nil
	}
	return a.ID, nil
}

func (a Author) GetById(id uint) (item interface{}, err error) {
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

func (a Author) GetAll() (items interface{}, err error) {
	var authors []Author
	tx := db.Find(&authors)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = authors
	return
}


func (a Author) Update() (err error) {
	tx := db.Save(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (a Author) Remove() (err error) {
	tx := db.Delete(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

