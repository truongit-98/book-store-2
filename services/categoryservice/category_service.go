package categoryservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"github.com/prometheus/common/log"
)

func GetAllWithBookCount() (interface{}, error){
	category := models.Category{}
	items, err := category.GetAllWithCount()
	if err != nil {
		log.Info(err.Error())
		return nil, responses.ErrSystem
	}
	collections := make(map[uint]map[string]interface{})
	for _, item := range items.([]responses.CategoryResponse) {
		if value, ok := collections[item.Id]; !ok {
			value = make(map[string]interface{})
			collections[item.Id] = value
			value["id"] = item.Id
			value["name"] = item.CategoryName
			value["topics"] = make([]map[string]interface{},0)
			topics := map[string]interface{}{
				"topic_id": item.TopicId,
				"topic_name": item.TopicName,
				"count": item.TotalBook,
			}
			value["topics"] = append(value["topics"].([]map[string]interface{}), topics)
			value["book_count"] = item.TotalBook
		} else {
			topics := map[string]interface{}{
				"topic_id": item.TopicId,
				"topic_name": item.TopicName,
				"count": item.TotalBook,
			}
			value["topics"] = append(value["topics"].([]map[string]interface{}), topics)
		}
	}
	return collections, nil
}

func CreateCategory(body requestbody.CategoryBody) error {
	category := models.Category{
		CategoryName: body.CategoryName,
	}
	err := models.Create(category)
	if err != nil {
		log.Info(err.Error())
		return err
	}
	return nil
}

func UpdateCategory(body requestbody.CategoryPutBody) error {
	category := models.Category{ID: body.ID}
	if exist, err := models.Exists(category); err != nil {
		log.Info(err.Error())
		return err
	} else if exist {
		category.CategoryName = body.CategoryName
		err = models.Update(category)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return responses.NotExisted
}

func DeleteCategory(id uint) error {
	category := models.Category{ID: id}
	if exist, err := models.Exists(category); err != nil {
		log.Info(err.Error())
		return err
	} else if exist {
		err = models.Remove(category)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return responses.NotExisted
}
