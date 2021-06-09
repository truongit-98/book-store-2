package bookservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"github.com/prometheus/common/log"
)

func CreateBookDetail(body requestbody.BookDetailPostRequestBody) (*models.BookDetail, error) {
	bDetail := &models.BookDetail{
		Price: body.PriceCover,
		PriceCover: body.PriceCover,
		NumberOfPage: body.NumberOfPage,
		Height: body.Height,
		Width: body.Width,
		Description: body.Description,
		Language: body.Language,
		BookTypeID: body.BookTypeID,
		PublisherID: body.PublisherID,
		FormatID: body.FormatID,
		TopicID: body.TopicID,
	}
	err := models.Create(bDetail)
	if err != nil {
		log.Info(err.Error())
		return nil, err
	}
	return bDetail, nil
}

