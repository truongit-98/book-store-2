 
package models

import (
	"BookStore/restapi/responses"
	"errors"
	"gorm.io/gorm"
)

type Book struct {
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ISBN string `gorm:"size:20" json:"isbn"`
	Title string `gorm:"size:50" json:"title"`
	PublisherYear int `json:"publisher_year"`
	Status         string  `gorm:"size:50" json:"status"`
	CoverImage string `gorm:"size:256" json:"cover_image"`
	Amount	int	`json:"amount"`
	NumberOfAccess int `json:"number_of_access"`
	SKU	string `gorm:"size:20" json:"sku"`
	Ratting float64 `json:"ratting"`
	CreatedAt int64 `json:"created_at"`
	BookTypeID *uint `json:"book_type_id"`
	Comments []*Comment `json:"comments"`
	OrderDetails []*OrderDetail `json:"order_details"`
}

func (b Book) GetPaginate(pos, count int32) (items interface{}, totalCount int32, err error) {
	panic("implement me")
}

func (b Book) GetWithFilter(categoryId, topicId, manufacturerId int32, priceMin, priceMax float64, sort string, offset, limit int32) (items interface{}, err error) {
	result := []responses.ProductResponse{}
	switch sort {
	case "popularity":
		tx := db.Raw("exec USP_GetBooksPopularWithFilter @categoryId = ?, @topicId = ?, @publisherId = ?, @priceMin = ?, @priceMax = ?, @offset = ?, @limit = ?", categoryId, topicId, manufacturerId, priceMin, priceMax, offset, limit).Scan(&result)
			if err = tx.Error; err != nil {
			return nil, err
		}
		break
	case "rating":
		tx := db.Raw("exec USP_GetBooksRatingWithFilter @categoryId = ?, @topicId = ?, @publisherId = ?, @priceMin = ?, @priceMax = ?, @offset = ?, @limit = ?", categoryId, topicId, manufacturerId, priceMin, priceMax, offset, limit).Scan(&result)
			if err = tx.Error; err != nil {
			return nil, err
		}
		break
	case "date":
		tx := db.Raw("exec USP_GetBooksNewNessWithFilter @categoryId = ?, @topicId = ?, @publisherId = ?, @priceMin = ?, @priceMax = ?, @offset = ?, @limit = ?", categoryId, topicId, manufacturerId, priceMin, priceMax, offset, limit).Scan(&result)
			if err = tx.Error; err != nil {
			return nil, err
		}
		break
	case "price":
		tx := db.Raw("exec USP_GetBooksLowToHighWithFilter @categoryId = ?, @topicId = ?, @publisherId = ?, @priceMin = ?, @priceMax = ?, @offset = ?, @limit = ?", categoryId, topicId, manufacturerId, priceMin, priceMax, offset, limit).Scan(&result)
			if err = tx.Error; err != nil {
			return nil, err
		}
		break
	case "price-desc":
		tx := db.Raw("exec USP_GetBooksHighToLowWithFilter @categoryId = ?, @topicId = ?, @publisherId = ?, @priceMin = ?, @priceMax = ?, @offset = ?, @limit = ?", categoryId, topicId, manufacturerId, priceMin, priceMax, offset, limit).Scan(&result)
			if err = tx.Error; err != nil {
			return nil, err
		}
		break
	default: 
		tx := db.Raw("exec USP_GetBooksPopularWithFilter @categoryId = ?, @topicId = ?, @publisherId = ?, @priceMin = ?, @priceMax = ?, @offset = ?, @limit = ?", categoryId, topicId, manufacturerId, priceMin, priceMax, offset, limit).Scan(&result)
			if err = tx.Error; err != nil {
			return nil, err
		}
		break
	}
	return result, nil
}
func (b Book) Exists() (exist bool, err error) {
	tx := db.First(&b, b.ID)
	if err := tx.Error;  err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return false, err
		}
		return false, nil
	}
	return true, nil
}

func (b Book) GetBySort(sort string) (items interface{}, err error) {
	var books []Book
	tx := db.Order(sort).Find(&books)
	if err = tx.Error; err != nil {
		return
	}
	items = books
	return
}

func (b Book) GetFeatured() (items interface{}, err error) {
	result := []responses.ProductResponse{}
	tx := db.Raw("EXEC USP_GetFeaturedProducts").Scan(&result)
	if err = tx.Error; err != nil {
		return
	}
	return result, nil 
}

func (b Book) GetBestSellers() (items interface{}, err error) {
	result := []responses.ProductResponse{}
	tx := db.Raw("EXEC USP_GetBestSellers").Scan(&result)
	if err = tx.Error; err != nil {
		return
	}
	return result, nil 

}

func (b Book) GetNewProducts() (items interface{}, err error) {
	result := []responses.ProductResponse{}
	tx := db.Raw("EXEC USP_GetNewProducts").Scan(&result)
	if err = tx.Error; err != nil {
		return
	}
	return result, nil 

}


func (b Book) GetWithConditions(conds interface{}, params ...interface{}) (items interface{}, err error) {
	var books []Book
	tx := db.Where(conds, params).Find(&books)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = books
	return
}
func (b Book) Create() (err error) {
	return nil
}
func (b Book) Create2() (id uint, err error) {
	tx := db.Where("title = ?", b.Title).First(&b)
	if errors.Is(tx.Error, gorm.ErrRecordNotFound){
		tx := db.Create(&b)
		if err = tx.Error; err != nil {
			return
		}
		return b.ID, nil
	}
	return b.ID, nil
}
func (b Book) GetById(id uint) (item interface{}, err error) {
	tx := db.First(&b, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound){
			err = responses.NotExisted
			return
		}
		err = responses.ErrSystem
		return
	}
	item = b
	return
}

func (b Book)GetBookInfo(id uint) (interface{}, error) {
	result := responses.ProductResponse{}
	tx := db.Raw("exec USP_GetInfoProduct @id = ?", id).Find(&result)
	if err := tx.Error; err != nil {
		return nil, err
	}
	return &result, nil
}

func (b Book) GetAll() (items interface{}, err error) {
	var books []Book
	tx := db.Preload("Comments").Find(&books)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	items = books
	return
}


func (b Book) Update() (err error) {
	tx := db.Save(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}

func (b Book) Remove() (err error) {
	tx := db.Delete(&b)
	if tx.Error != nil {
		err = responses.ErrSystem
		return
	}
	return
}


