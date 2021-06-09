package bookservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"github.com/prometheus/common/log"
)

func GetAll() ([]models.Book, error){
	book := models.Book{}
	items, err := models.GetAll(book)
	if err != nil {
		log.Info(err.Error())
		return make([]models.Book, 0), err
	}
	return items.([]models.Book), nil
}

func GetFeaturedProducts() (interface{}, error){
	data, err := (models.Book{}).GetFeatured()
	if err != nil {
		log.Info(err.Error())
		return nil, responses.ErrSystem
	}
	return data, nil
}

func GetBestSellers() (interface{}, error){
	data, err := (models.Book{}).GetBestSellers()
	if err != nil {
		log.Info(err.Error())
		return nil, responses.ErrSystem
	}
	return data, nil
}

func GetNewProducts() (interface{}, error){
	data, err := (models.Book{}).GetNewProducts()
	if err != nil {
		log.Info(err.Error())
		return nil, responses.ErrSystem
	}
	log.Info(data)
	return data, nil
}


func CreateBook(body requestbody.BookPostRequestBody) error {
	book := models.Book{
		ISBN: body.ISBN,
		SKU: body.SKU,
		Title: body.Title,
		CoverImage: body.CoverImage,
		Amount: int(body.Amount),
		BookTypeID: &body.TypeID,
	}
	err := models.Create(book)
	if err != nil {
		log.Info(err.Error())
		return err
	}
	return nil
}

func UpdateBook(body requestbody.BookPutRequestBody) error {
	book := models.Book{ID: body.ID}
	if exist, err := models.Exists(book); err != nil {
		log.Info(err.Error())
		return err
	} else if exist {
		book := models.Book{
			ID: body.ID,
			ISBN: body.ISBN,
			SKU: body.SKU,
			Title: body.Title,
			CoverImage: body.CoverImage,
			Amount: int(body.Amount),
			BookTypeID: &body.BookTypeID,
		}
		err = models.Update(book)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return responses.NotExisted
}

func DeleteBook(id uint) error {
	book := models.Book{ID: id}
	if exist, err := models.Exists(book); err != nil {
		log.Info(err.Error())
		return err
	} else if exist {
		err = models.Remove(book)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return responses.NotExisted
}

func GetBookInfo(id uint) (interface{}, error) {
	book, err := (models.Book{}).GetBookInfo(id)
	if err != nil {
		log.Info(err.Error())
		return nil, err
	}
	return book, nil
}

func GetWithFilter(categoryId, topicId, manufacturerId int32, priceMin, priceMax float64, sort string, offset, limit int32) (interface{}, error)  {

	if topicId != 0 {
		categoryId = 0
	}
	if offset < 0 {
		offset = 0
	} 
	if limit <= 0 {
		limit = 10
	}
	if offset > limit {
		offset = limit
	}

	if priceMax <= 0 {
		priceMax = 10000000
	}

	items, err := (&models.Book{}).GetWithFilter(categoryId, topicId, manufacturerId, priceMin, priceMax, sort, offset, limit)
	if err != nil {
		log.Info(err)
		return nil, responses.ErrSystem
	}
	return items, nil

}
