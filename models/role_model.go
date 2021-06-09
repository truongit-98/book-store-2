package models

import (
	"errors"
	"gorm.io/gorm"
)

type Role struct {
	ID        uint               `json:"id" xml:"id"`
	Name      string             `json:"name" xml:"name"`
	CreatedAt int64              `json:"created_at" xml:"created_at"`
	UpdatedAt int64              `json:"updated_at" xml:"updated_at"`
	RolePerms []*RolePermissions `json:"role_perms" xml:"role_perms"`
	AdminRoles []*AdminRole `json:"admin_roles" xml:"admin_roles"`
}

func (r* Role) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	var Roles []Role
	tx := db.Offset(int(pos)).Limit(int(count)).Preload("Admins").Preload("RolePerms.Controls").Find(&Roles)

	if err = tx.Error; err != nil {
		return
	}
	items = Roles
	return
}

func (r* Role) Exists() (bool, error) {
	tx := db.First(&r, r.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (r* Role) GetBySort(sort string) (items interface{}, err error) {
	var Roles []Role
	tx := db.Order(sort).Find(&Roles)
	if err = tx.Error; err != nil {
		return
	}
	items = Roles
	return
}

func (r* Role) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var Roles []Role
	tx := db.Preload("RolePerms.Controls").Where(conds, params).Find(&Roles)
	if err = tx.Error; err != nil {
		return
	}
	items = Roles
	return
}

func (r* Role) Create() (err error) {
	tx := db.Create(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (r* Role)GetAllWithPreload() ([]Role, error) {
	var Roles = make([]Role,0)
	tx := db.Preload("RolePerms.RolePermControls").Find(&Roles)
	if err := tx.Error; err != nil {
		return Roles, err
	}
	return Roles, nil
}

func (r* Role) GetById(id uint) (item interface{}, err error) {
	tx := db.Preload("RolePerms.RolePermControls").First(&r, id)
	if err = tx.Error; err != nil {
		return
	}
	item = r
	return
}


func (r* Role) GetAll() (items interface{}, err error) {
	var Roles []Role
	tx := db.Find(&Roles)
	if err = tx.Error; err != nil {
		return
	}
	items = Roles
	return
}

func (r* Role) Update() (err error) {
	tx := db.Save(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}

func (r* Role) Remove() (err error) {
	tx := db.Delete(&r)
	if err = tx.Error; err != nil {
		return
	}
	return
}
