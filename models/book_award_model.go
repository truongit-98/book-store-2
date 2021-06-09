package models

import (
	"BookStore/restapi/responses"
	"errors"
)

type BookTypeAward struct {
	AwardID uint `json:"award_id"`
	BookTypeID uint `json:"book_type_id"`
	AwardYear uint8 `json:"award_year"`
}

func (b BookTypeAward) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (b BookTypeAward) Exists() (exist bool, err error) {
	//tx := db.Where("book_id = ? AND award_id = ?", b.BookID, b.AwardID).First(&b)
	//if err := tx.Error;  err != nil {
	//	if !errors.Is(err, gorm.ErrRecordNotFound) {
	//		return false, err
	//	}
	//	return false, nil
	//}
	return true, nil
}

func (b BookTypeAward) GetBySort(sort string) (items interface{}, err error) {
	var bookAwards []BookTypeAward
	tx := db.Order(sort).Find(&bookAwards)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookAwards
	return
}

func (b BookTypeAward) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var bookAwards []BookTypeAward
	tx := db.Where(conds, params).Find(&bookAwards)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookAwards
	return
}

func (b BookTypeAward) Create() (err error) {
	tx := db.Create(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (b BookTypeAward) GetById(id uint) (item interface{}, err error) {
	return nil, errors.New("not supported method")
}

func (b BookTypeAward) GetAll() (items interface{}, err error) {
	var bookAwards []BookTypeAward
	tx := db.Find(&bookAwards)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = bookAwards
	return
}

func (b BookTypeAward) Update() (err error) {
	tx := db.Save(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (b BookTypeAward) Remove() (err error) {
	tx := db.Delete(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}


