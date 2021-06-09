package customerservice

import (
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/responses"
	"errors"
	"github.com/prometheus/common/log"
	"gorm.io/gorm"
)

func GetAll() ([]models.Customer, error){
	admin := models.Customer{}
	items, err := models.GetAll(&admin)
	if err != nil {
		log.Info(err.Error())
		return make([]models.Customer, 0), err
	}
	return items.([]models.Customer), nil
}
func GetPaginate(pos, count int32)([]models.Customer, int32, error)  {
	result := make([]models.Customer, 0)
	data, totalCount, err := models.GetPaginate(&models.Customer{}, pos, count)
	if err != nil {
		log.Info(err.Error())
		return result, totalCount,  responses.ErrSystem
	}
	result = data.([]models.Customer)
	return result, totalCount, nil
}

func RegisterUserAccount(body requestbody.AccountRequest) error {
	conds := "email = ?"
	data, err := models.GetWithCondition(&models.Customer{}, conds, body.Email)
	if err != nil {
		log.Info(err.Error())
		return err
	}
	if len(data.([]models.Customer)) <= 0 {
		admin := models.Customer{
			Email: body.Email,
			Password: body.Password,
			FullName: body.FullName,
			Address: body.Address,
			Phone: body.Phone,
		}
		err := models.Create(&admin)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return errors.New("email is already taken")
}

func LoginUser(body requestbody.AccountLoginRequest) (int, error) {
	uid, err := (&models.Customer{}).Login(body.Email, body.Password)
	if err != nil {
		log.Info(err.Error())
		return 0, responses.ErrSystem
	}
	if uid <= 0 {
		return -1, responses.NotExisted
	}
	return uid, nil
}

func UpdateUserAccount(body requestbody.AccountPutRequest) error {
	admin := models.Admin{ID: body.Id}
	if exist, err := models.Exists(&admin); err != nil {
		log.Info(err.Error())
		return err
	} else if exist {
		newAdmin := models.Customer{
			FullName: body.FullName,
			Address: body.Address,
			Phone: body.Phone,
		}
		err = models.Update(&newAdmin)
		if err != nil {
			log.Info(err.Error())
			return err
		}
		return nil
	}
	return responses.NotExisted
}

func GetCustomerById(custId int) (*models.Customer, error) {
	customer, err := models.GetById(&models.Customer{}, uint(custId))
	if err != nil{
		log.Info(err)
		if errors.Is(gorm.ErrRecordNotFound, err){
			return nil, responses.NotExisted
		}
		return nil, responses.ErrSystem
	}
	return customer.(*models.Customer), nil
}

func GetCustomerInfo(custId int32) (*models.Customer, error) {
	customer, err := (&models.Customer{}).GetCustomerInfo(uint(custId))
	if err != nil{
		log.Info(err)
		if errors.Is(gorm.ErrRecordNotFound, err){
			return nil, responses.NotExisted
		}
		return nil, responses.ErrSystem
	}
	return customer.(*models.Customer), nil
}

func SearchCustomer(name string) ([]models.Customer, error) {
	data, err := (&models.Customer{}).SearchByName(name)
	if err != nil {
		log.Info(err)
		return make([]models.Customer, 0), responses.ErrSystem
	}
	return data.([]models.Customer), nil
}
