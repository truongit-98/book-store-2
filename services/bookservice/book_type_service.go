package bookservice

import (
	"BookStore/consts"
	"BookStore/models"
	"BookStore/requestbody"
	"github.com/prometheus/common/log"
)

func CreateBookType(body requestbody.BookTypePostRequestBody) (*models.BookType, error) {
	var (
		typeStr string
		ok bool
	)

	if typeStr, ok = consts.MAP_BOOK_TYPE[body.BookType]; !ok {
		typeStr = consts.MAP_BOOK_TYPE[1]
	}
	bType := &models.BookType{
		Type: typeStr,
		TypeName: body.TypeName,
		Episodes: body.Episodes,
	}
	err := models.Create(bType)
	if err != nil {
		log.Info(err.Error())
		return nil, err
	}
	return bType, nil
}
