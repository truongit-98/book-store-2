package models

import (
	"BookStore/restapi/responses"
	"errors"
)

type BookTypeAuthor struct {
	BookTypeID   uint `json:"book_type_id"`
	AuthorID uint     `json:"author_id"`
	Role     string `gorm:"size:50;" json:"role"`
}

func (b BookTypeAuthor) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (b BookTypeAuthor) Exists() (exist bool, err error) {
	//tx := db.Where("book_id = ? AND author_id = ?", b.BookTypeID, b.AuthorID).First(&b)
	//if err := tx.Error;  err != nil {
	//	if !errors.Is(err, gorm.ErrRecordNotFound) {
	//		return false, err
	//	}
	//	return false, nil
	//}
	return true, nil
}

func (b BookTypeAuthor) GetBySort(sort string) (items interface{}, err error) {
	var bookAuthors []BookTypeAuthor
	tx := db.Order(sort).Find(&bookAuthors)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookAuthors
	return
}

func (b BookTypeAuthor) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var bookAuthors []BookTypeAuthor
	tx := db.Where(conds, params).Find(&bookAuthors)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookAuthors
	return
}

func (b BookTypeAuthor) Create() (err error) {
	tx := db.Create(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (b BookTypeAuthor) GetById(id uint) (item interface{}, err error) {
	return nil, errors.New("not supported method")
}

func (b BookTypeAuthor) GetAll() (items interface{}, err error) {
	var bookAuthors []BookTypeAuthor
	tx := db.Find(&bookAuthors)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookAuthors
	return
}

func (b BookTypeAuthor) Update() (err error) {
	tx := db.Save(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (b BookTypeAuthor) Remove() (err error) {
	tx := db.Delete(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}
