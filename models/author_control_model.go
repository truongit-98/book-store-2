package models

import (
	"errors"
	"gorm.io/gorm"
)

type AuthorControl struct {
	ID uint `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
	Description string `json:"description"`
	Action string `json:"action" xml:"action"`
	CreatedAt int64 `json:"created_at" xml:"created_at"`
	UpdatedAt int64 `json:"updated_at" xml:"updated_at"`
	RolePermControls []*RolePermissionControl `json:"controls" xml:"controls"`
}

func (a AuthorControl) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	var controls []AuthorControl
	tx := db.Offset(int(pos)).Limit(int(count)).Find(&controls)
	if err = tx.Error; err != nil {
		return
	}
	items = controls
	return
}

func (a AuthorControl) Exists() (exist bool, err error) {
	tx := db.First(&a, a.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (a AuthorControl) GetBySort(sort string) (items interface{}, err error) {
	var controls []AuthorControl
	tx := db.Order(sort).Find(&controls)
	if err = tx.Error; err != nil {
		return false, err
	}
	items = controls
	return
}

func (a AuthorControl) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var controls []AuthorControl
	tx := db.Where(conds, params).Find(&controls)
	if err = tx.Error; err != nil {
		return
	}
	items = controls
	return
}

func (a AuthorControl) Create() (err error) {
	tx := db.Create(&a)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (a AuthorControl) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&a, id)
	if err = tx.Error; err != nil {
		return
	}
	item = a
	return
}

func (a AuthorControl) GetAll() (items interface{}, err error) {
	var controls []AuthorControl
	tx := db.Find(&controls)
	if err = tx.Error; err != nil {
		return
	}
	items = controls
	return
}


func (a AuthorControl) Update() (err error) {
	tx := db.Save(&a)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (a AuthorControl) Remove() (err error) {
	tx := db.Delete(&a)
	if err = tx.Error; err != nil {
		return
	}
	return
}
