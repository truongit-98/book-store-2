package models

import (
	"errors"
	"gorm.io/gorm"
)

type AdminRole struct {
	ID uint `json:"id" xml:"id"`
	AdminID   uint         `json:"admin_id" xml:"admin_id"`
	RoleID    uint         `json:"role_id" xml:"role_id"`
	CreatedAt int64 `json:"created_at" xml:"created_at"`
}

func (r* AdminRole) Exists() (exist bool, err error) {
	tx := db.Where("admin_id = ? AND role_id = ?", r.AdminID, r.RoleID).First(&r)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r* AdminRole) GetBySort(sort string) (items interface{}, err error) {
	var adminRoles []AdminRole
	tx := db.Order(sort).Find(&adminRoles)
	if err = tx.Error; err != nil {
		return
	}
	items = adminRoles
	return
}
func (r* AdminRole) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	var adminRoles []AdminRole
	tx := db.Offset(int(pos)).Limit(int(count)).Find(&adminRoles)
	if err = tx.Error; err != nil {
		return
	}
	items = adminRoles
	return
}

func (r* AdminRole) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var adminRoles []AdminRole
	tx := db.Where(conds, params).Find(&adminRoles)
	if err = tx.Error; err != nil {
		return
	}
	items = adminRoles
	return
}

func (r* AdminRole) Create() (err error) {
	tx := db.Create(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (r* AdminRole) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&r, id)
	if err = tx.Error; err != nil {
		return
	}
	item = r
	return
}

func (r* AdminRole) GetAll() (items interface{}, err error) {
	var adminRoles []AdminRole
	tx := db.Find(&adminRoles)
	if err = tx.Error; err != nil {
		return
	}
	items = adminRoles
	return
}

func (r* AdminRole) Update() (err error) {
	return nil
}

func (r* AdminRole) Remove() (err error) {
	tx := db.Delete(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}
