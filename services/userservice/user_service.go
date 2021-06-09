package userservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"errors"
	"github.com/prometheus/common/log"
)

func GetAll() ([]models.Customer, error){
	book := models.Customer{}
	items, err := models.GetAll(book)
	if err != nil {
		log.Info(err.Error())
		return make([]models.Customer, 0), err
	}
	return items.([]models.Customer), nil
}
func GetPaginate(pos, count int32)([]models.Customer, int32, error)  {
	result := make([]models.Customer, 0)
	data, totalCount, err := models.GetPaginate(models.Customer{}, pos, count)
	if err != nil {
		log.Info(err.Error())
		return result, totalCount,  responses.ErrSystem
	}
	result = data.([]models.Customer)
	return result, totalCount, nil
}

func RegisterUserAccount(body requestbody.UserRequestBody) error {
	conds := "email = ?"
	data, err := models.GetWithCondition(models.Customer{}, conds, body.Email)
	if err != nil {
		log.Info(err.Error())
		return err
	}
	if len(data.([]models.Customer)) <= 0 {
		user := models.Customer{
			Email: body.Email,
			Password: body.Password,
			FullName: body.FullName,
			Address: body.Address,
			Phone: body.Phone,
		}
		err := models.Create(user)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return errors.New("email is already taken")
}

func LoginUser(body requestbody.UserLoginBody) (int, error) {
	uid, err := models.Customer{}.Login(body.Email, body.Password)
	if err != nil {
		log.Info(err.Error())
		return 0, responses.ErrSystem
	}
	if uid == 0 {
		return 0, errors.New("incorrect username or password")
	}
	return uid, nil
}

// func UpdateuserAccount(body requestbody.CustomerPutRequest) error {
// 	user := models.Customer{ID: body.Id}
// 	if exist, err := models.Exists(user); err != nil {
// 		log.Info(err.Error())
// 		return err
// 	} else if exist {
// 		newuser := models.Customer{
// 			FullName: body.FullName,
// 			Address: body.Address,
// 			Phone: body.Phone,
// 		}
// 		err = models.Update(newuser)
// 		if err != nil {
// 			log.Info(err.Error())
// 			return err
// 		}
// 		return nil
// 	}
// 	return responses.NotExisted
// }

//func DeleteBook(id uint) error {
//	book := models.Book{ID: id}
//	if exist, err := models.Exists(book); err != nil {
//		log.Info(err.Error())
//		return err
//	} else if exist {
//		err = models.Remove(book)
//		if err != nil {
//			log.Info(err.Error())
//			return err
//		}
//		return nil
//	}
//	return responses.NotExisted
//}
