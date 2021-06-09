package models

import (
	"errors"
	"gorm.io/gorm"
)

type RolePermissionControl struct {
	ID uint `json:"id" xml:"id"`
	RolePermissionsID uint  `json:"role_permissions_id" xml:"role_permissions_id"`
	AuthorControlID  uint  `json:"author_control_id" xml:"author_control_id"`
	CreatedAt        int64 `json:"created_at" xml:"created_at"`
}

func (r* RolePermissionControl) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	var routers []RolePermissionControl
	tx := db.Offset(int(pos)).Limit(int(count)).Find(&routers)
	if err = tx.Error; err != nil {
		return
	}
	items = routers
	return
}

func (r* RolePermissionControl) Exists() (exist bool, err error) {
	tx := db.Where("author_control_id = ? AND role_permissions_id = ?", r.AuthorControlID, r.RolePermissionsID).First(&r)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r* RolePermissionControl) GetBySort(sort string) (items interface{}, err error) {
	var rolePerms []RolePermissionControl
	tx := db.Order(sort).Find(&rolePerms)
	if err = tx.Error; err != nil {
		return
	}
	items = rolePerms
	return
}

func (r* RolePermissionControl) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var rolePerms []RolePermissionControl
	tx := db.Where(conds, params).Find(&rolePerms)
	if err = tx.Error; err != nil {
		return
	}
	items = rolePerms
	return
}

func (r* RolePermissionControl) Create() (err error) {
	tx := db.Create(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (r* RolePermissionControl) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&r, id)
	if err = tx.Error; err != nil {
		return
	}
	item = r
	return
}

func (r* RolePermissionControl) GetAll() (items interface{}, err error) {
	var rolePerms []RolePermissionControl
	tx := db.Find(&rolePerms)
	if err = tx.Error; err != nil {
		return
	}
	items = rolePerms
	return
}

func (r* RolePermissionControl) Update() (err error) {
	return nil
}

func (r* RolePermissionControl) Remove() (err error) {
	tx := db.Delete(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}