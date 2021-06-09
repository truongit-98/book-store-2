package models

import (
	"errors"
	"gorm.io/gorm"
)

type Permission struct {
	ID        uint               `json:"id" xml:"id"`
	Name      string             `json:"name" xml:"name"`
	Path      string             `json:"path" xml:"path"`
	CreatedAt int64              `json:"created_at" xml:"created_at"`
	UpdatedAt int64              `json:"updated_at" xml:"updated_at"`
	RolePerms []*RolePermissions `json:"role_perms" xml:"role_perms"`
}

func (p Permission) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	var perms []Permission
	tx := db.Offset(int(pos)).Limit(int(count)).Find(&perms)
	if err = tx.Error; err != nil {
		return
	}
	items = perms
	return
}

func (p Permission) Exists() (exist bool, err error) {
	tx := db.First(&p, p.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (p Permission) GetBySort(sort string) (items interface{}, err error) {
	var Permissions []Permission
	tx := db.Order(sort).Find(&Permissions)
	if err = tx.Error; err != nil {
		return
	}
	items = Permissions
	return
}

func (p Permission) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var Permissions []Permission
	tx := db.Where(conds, params).Find(&Permissions)
	if err = tx.Error; err != nil {
		return
	}
	items = Permissions
	return}

func (p Permission) Create() (err error) {
	tx := db.Create(&p)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p Permission) GetById(id uint) (item interface{}, err error) {
	tx := db.Preload("RolePerms.RolePermControls").First(&p, id)
	if err = tx.Error; err != nil {
		return
	}
	item = p
	return
}

func (p Permission) GetAll() (items interface{}, err error) {
	var Permissions []Permission
	tx := db.Find(&Permissions)
	if err = tx.Error; err != nil {
		return
	}
	items = Permissions
	return
}

func (p Permission) Update() (err error) {
	tx := db.Save(&p)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (p Permission) Remove() (err error) {
	tx := db.Delete(&p)
	if err = tx.Error; err != nil {
		return
	}
	return
}
