package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Customer struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Email string `gorm:"size:50;not null" json:"email"`
	Password string `gorm:"size:20;not null" json:"password"`
	FullName string `gorm:"size:50" json:"full_name"`
	Address string `gorm:"size:256" json:"address"`
	Phone string `gorm:"size:10" json:"phone"`
	Activate bool `json:"activate"`
	Avatar string `gorm:"size:256" json:"avatar"`
	Orders []*Order `json:"orders"`
	Comments []*Comment `json:"comments"`
}

func (c* Customer) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	var cus = make([]Customer, 0)
	err = db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Offset(int(pos)).Limit(int(count)).Find(&cus)
		if err = tx.Error; err != nil {
			return err
		}
		items = cus
		var total int
		if err = tx.Raw("select count(*) from customers").Scan(&total).Error; err != nil {
			return err
		}
		totalCount = int32(total)
		return nil
	})
	if err != nil {
		return nil, 0, err
	}
	return
}
func (c* Customer) Login(email, password string) (int, error) {
	var uid int
	tx := db.Raw("SELECT dbo.FUNC_LoginUser(?, ?)", email, password).Scan(&uid)
	if err := tx.Error; err != nil {
		return 0, err
	}
	return uid , nil
}
func (c* Customer) Exists() (exist bool, err error) {
	tx := db.First(&c, c.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (c* Customer) GetBySort(sort string) (items interface{}, err error) {
	var customers []Customer
	tx := db.Order(sort).Find(&customers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = customers
	return
}

func (c* Customer) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var customers []Customer
	tx := db.Where(conds, params).Find(&customers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = customers
	return}

func (c* Customer) Create() (err error) {
	tx := db.Create(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (c* Customer) GetCustomerInfo(id uint) (item interface{}, err error) {
	tx := db.Preload("Orders").Preload("Comments").First(&c, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = c
	return
}

func (c* Customer) SearchByName(name string) (items interface{}, err error) {
	var customers []Customer
	tx := db.Raw("exec usp_searchUserByName @name = ?", name).Scan(&customers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = customers
	return
}

func (c* Customer) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&c, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = c
	return
}

func (c* Customer) GetAll() (items interface{}, err error) {
	var customers []Customer
	tx := db.Find(&customers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = customers
	return
}

func (c* Customer) Update() (err error) {
	tx := db.Save(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (c* Customer) Remove() (err error) {
	tx := db.Delete(&c)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}
