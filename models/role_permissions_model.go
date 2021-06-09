package models

import (
	"errors"
	"gorm.io/gorm"
)

type RolePermissions struct {
	ID uint `json:"id" xml:"id"`
	CreatedAt int64 `json:"created_at" xml:"created_at"`
	UpdatedAt int64 `json:"updated_at" xml:"updated_at"`
	RoleID uint `json:"role_id" xml:"role_id"`
	PermissionID uint `json:"permission_id" xml:"permission_id"`
	RolePermControls []*RolePermissionControl `json:"controls" xml:"controls"`
}

func (r* RolePermissions) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	var rolePerms []RolePermissions
	tx := db.Offset(int(pos)).Limit(int(count)).Find(&rolePerms)
	if err = tx.Error; err != nil {
		return
	}
	items = rolePerms
	return
}

func (r* RolePermissions) Exists() (exist bool, err error) {
	tx := db.First(&r, r.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r* RolePermissions) GetBySort(sort string) (items interface{}, err error) {
	var rolePerms []RolePermissions
	tx := db.Order(sort).Find(&rolePerms)
	if err = tx.Error; err != nil {
		return
	}
	items = rolePerms
	return
}

func (r* RolePermissions) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var rolePerms []RolePermissions
	tx := db.Where(conds, params).Find(&rolePerms)
	if err = tx.Error; err != nil {
		return
	}
	items = rolePerms
	return
}

func (r* RolePermissions) Create() (err error) {
	tx := db.Create(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (r* RolePermissions) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&r, id)
	if err = tx.Error; err != nil {
		return
	}
	item = r
	return
}

func (r* RolePermissions) GetAll() (items interface{}, err error) {
	var rolePerms []RolePermissions
	tx := db.Find(&rolePerms)
	if err = tx.Error; err != nil {
		return
	}
	items = rolePerms
	return
}

func (r* RolePermissions) Update() (err error) {
	tx := db.Save(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (r* RolePermissions) Remove() (err error) {
	tx := db.Delete(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}