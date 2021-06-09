package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Voucher struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Value float64 `json:"value"`
	Code string `gorm:"size:50;unique" json:"code"`
	Expiry int64 `json:"expiry"`
	OrderVouchers []*OrderVouchers `json:"order_vouchers"`
}

func (v* Voucher) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (v* Voucher) Exists() (exist bool, err error) {
	tx := db.First(&v, v.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (v* Voucher) GetBySort(sort string) (items interface{}, err error) {
	var vouchers []Voucher
	tx := db.Order(sort).Find(&vouchers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = vouchers
	return
}

func (v* Voucher) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var vouchers []Voucher
	tx := db.Where(conds, params).Find(&vouchers)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = vouchers
	return
}

func (v* Voucher) Create() (err error) {
	tx := db.Create(&v)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (v* Voucher) GetById(id uint) (item interface{}, err error) {
	tx := db.First(v, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = v
	return
}

func (v* Voucher) GetByCode(code string) (interface{}, error) {
	err := db.Where("code = ?", code).First(&v).Error
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (v* Voucher) GetAll() (items interface{}, err error) {
	var topics []Topic
	tx := db.Find(&topics)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = topics
	return
}

func (v* Voucher) Update() (err error) {
	tx := db.Save(&v)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (v* Voucher) Remove() (err error) {
	tx := db.Delete(&v)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}


