package models

import (
	"github.com/prometheus/common/log"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"sync"
)

var (
	db *gorm.DB
	factoryIns RepoFactoryIf
	syncF sync.Once
)

func InitDB() (*gorm.DB, error) {

	var err error
	if db == nil {
		db, err = gorm.Open(sqlserver.New(sqlserver.Config{
			DSN: "sqlserver://sa:Truong123@localhost:1433?database=testDB",
		}), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		// dsn := "root:truong123456@tcp(127.0.0.1:3306)/testDB?charset=utf8mb4&parseTime=True&loc=Local"
		// db, err  = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 	SkipDefaultTransaction: true,
		// })
		if err != nil {
			log.Info(err.Error(), "models/global.go::13")
			return nil, err
		}
		db.AutoMigrate(&Admin{})
		db.AutoMigrate(&Author{})
		db.AutoMigrate(&Award{})
		db.AutoMigrate(&BookTypeAuthor{})
		db.AutoMigrate(&BookTypeAward{})
		db.AutoMigrate(&Book{})
		db.AutoMigrate(&BookType{})
		db.AutoMigrate(&BookDetail{})
		db.AutoMigrate(&Category{})
		db.AutoMigrate(&Comment{})
		db.AutoMigrate(&Customer{})
		db.AutoMigrate(&Delivery{})
		db.AutoMigrate(&Format{})
		db.AutoMigrate(&OrderDetail{})
		db.AutoMigrate(&Order{})
		db.AutoMigrate(&Payment{})
		db.AutoMigrate(&Publisher{})
		db.AutoMigrate(&Topic{})
		db.AutoMigrate(&Voucher{})
		db.AutoMigrate(&OrderVouchers{})
		db.AutoMigrate(&Role{})
		db.AutoMigrate(&Permission{})
		db.AutoMigrate(&Admin{})
		db.AutoMigrate(&AdminRole{})
		db.AutoMigrate(&RolePermissions{})
		db.AutoMigrate(&AuthorControl{})
		db.AutoMigrate(&RolePermissionControl{})
		db.AutoMigrate(&OrderVouchers{})
		//Create(&AuthorControl{Name: "View",  Action: "GET"})
		//Create(&AuthorControl{Name: "Add", Action: "POST"})
		//Create(&AuthorControl{Name: "Edit", Action: "EDIT"})
		//Create(&AuthorControl{Name: "Delete", Action: "DELETE"})

	}
	return db, nil
}

func InitService() RepoFactoryIf {
	syncF.Do(func() {
		factoryIns = NewFactoryInstance()
	})
	return factoryIns
}
