package models

import (
	"BookStore/restapi/data"
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type ReceiveInfo struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

type Order struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Total float64 `json:"total"`
	DateCreated int64 `json:"date_created"`
	Status string `gorm:"size:50" json:"status"`
	OrderDate int64 `json:"order_date"`
	ReceiveInfo *string `json:"receive_info" orm:"models.ReceiveInfo"`
	CustomerID *uint `json:"customer_id"`
	PaymentID *uint `json:"payment_id"`
	Creator *uint `json:"creator"`
	Delivery *uint `json:"delivery"`
	OrderDetails []*OrderDetail `json:"order_details"`
	OrderVouchers []*OrderVouchers `json:"order_vouchers"`
}

func (o* Order) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (o* Order) Exists() (exist bool, err error) {
	tx := db.First(&o, o.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (o* Order) GetBySort(sort string) (items interface{}, err error) {
	var orders []Order
	tx := db.Order(sort).Find(&orders)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = orders
	return
}

func (o *Order) GetWithFilter(status string, start, end, offset, limit int32) ([]data.OrderResponse, error) {
	var result = make([]data.OrderResponse,0)
	tx := db.Raw("exec [USP_GetOrderWithFilter] @status = ?, @startTime = ?, @endTime = ?, @offset = ?, @limit = ?", status, start, end, offset, limit).Scan(&result)
	if err := tx.Error; err != nil {
		return make([]data.OrderResponse, 0), err
	}
	return result, nil
}

func (o* Order) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var orders []Order
	tx := db.Where(conds, params).Find(&orders)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = orders
	return}

func (o* Order) Create() (err error) {
	tx := db.Create(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (o* Order) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&o, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = o
	return
}

func (o* Order) GetAll() (items interface{}, err error) {
	var orders []Order
	tx := db.Order("order_date desc").Find(&orders)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = orders
	return
}

func (o* Order) Update() (err error) {
	tx := db.Save(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (o* Order) Remove() (err error) {
	tx := db.Delete(&o)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}
