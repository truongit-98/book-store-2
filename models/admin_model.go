package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Admin struct {
	ID uint `gorm:"primaryKey;autoIncrement;" json:"id"`
	Email string `gorm:"size:50;not null;index;unique" json:"email"`
	Password string `gorm:"size:20;not null" json:"password"`
	FullName string `gorm:"size:50" json:"full_name"`
	Address string `gorm:"size:256" json:"address"`
	Phone string `gorm:"size:10" json:"phone"`
	Orders []*Order `gorm:"foreignKey:creator" json:"orders"`
	AdminRoles []*AdminRole `json:"admin_roles" xml:"admin_roles"`
}

func (a* Admin) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	var Admins []Admin
	tx := db.Offset(int(pos)).Limit(int(count)).Find(&Admins)
	if err = tx.Error; err != nil {
		return
	}
	items = Admins
	return
}

func (a* Admin) Login(email, password string) (int, error) {
	var result int
	tx := db.Raw("SELECT dbo.FUNC_LoginAdmin(?, ?)", email, password).Scan(&result)
	if err := tx.Error; err != nil {
		return -1, err
	}
	return result , nil
}

func (a* Admin) Exists() (exist bool, err error) {
	tx := db.First(&a, a.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (a* Admin) GetAdminInfo(adminId uint) (items interface{}, err error) {
	tx := db.Preload("AdminRoles").First(&a, adminId)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = a
	return
}
func (a* Admin) GetBySort(sort string) (items interface{}, err error) {
	var admins []Admin
	tx := db.Order(sort).Find(&admins)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = admins
	return
}

func (a* Admin) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var admins []Admin
	tx := db.Where(conds, params).Find(&admins)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = admins
	return
}

func (a* Admin) Create() (err error) {
	tx := db.Create(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (a* Admin) GetById(id uint) (item interface{}, err error) {
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

func (a* Admin) GetAll() (items interface{}, err error) {
	var admins []Admin
	tx := db.Find(&admins)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = admins
	return
}

func (a* Admin) Update() (err error) {
	tx := db.Save(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (a* Admin) Remove() (err error) {
	tx := db.Delete(&a)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

