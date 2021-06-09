package orderservice

import (
	"BookStore/events"
	"BookStore/models"
	"BookStore/requestbody"
	"BookStore/restapi/data"
	"BookStore/restapi/responses"
	"BookStore/services/commonservice"
	"BookStore/services/customerservice"
	"BookStore/utils"
	"encoding/json"
	"errors"
	"github.com/prometheus/common/log"
	"gorm.io/gorm"
	"strconv"
	"strings"
	"time"
)
type typeStatus = int
const (
	STATUS_PENDING typeStatus = iota + 1
	STATUS_PROCESSING
	STATUS_SUCCESS
)

var mapStatusOrder = map[int]string{
	STATUS_PENDING: "Đang chờ xử lý",
	STATUS_PROCESSING: "Đang xử lý",
	STATUS_SUCCESS: "Thành công",
}

func handleAddress(cityId, districtId, wardId string) (string, error) {
	citiesByte, err := commonservice.LoadCities()
	if err != nil {
		return "", err
	}
	var citiesMap map[string]data.Administrative
	err = json.Unmarshal(citiesByte, &citiesMap)
	if err != nil {
		log.Info(err)
		return "", responses.ErrSystem
	}
	if _, exist := citiesMap[cityId]; !exist {
		return "", responses.NotExisted
	}
	districtsByte, err := commonservice.LoadDistrictForCity(cityId)
	if err != nil {
		return "", err
	}
	var districtsMap map[string]data.Administrative
	err = json.Unmarshal(districtsByte, &districtsMap)
	if err != nil {
		log.Info(err)
		return "", responses.ErrSystem
	}
	if _, exist := districtsMap[districtId]; !exist {
		return "", responses.NotExisted
	}
	wardsByte, err := commonservice.LoadCities()
	if err != nil {
		return "", err
	}
	var wardsMap map[string]data.Administrative
	err = json.Unmarshal(wardsByte, &wardsMap)
	if err != nil {
		log.Info(err)
		return "", responses.ErrSystem
	}
	if ward, exist := citiesMap[cityId]; !exist {
		return "", responses.NotExisted
	} else {
		return ward.PathWithType, nil
	}
}

func CheckoutOrder(req requestbody.OrderInformation, user string) error {
	userId, err := strconv.Atoi(user)
	if err != nil {
		log.Info(err.Error())
		return responses.BadRequest
	}
	userInfo, err := customerservice.GetCustomerById(userId)
	if err != nil {
		return err
	}
	var receiveInfo models.ReceiveInfo
	if req.ReceiveInfo != nil {
		addressDetail := req.ReceiveInfo.AddressDetail
		tempAddress, err := handleAddress(req.ReceiveInfo.CityId, req.ReceiveInfo.DistrictId, req.ReceiveInfo.WardId)
		if err != nil {
			return err
		}
		address := strings.Join([]string{addressDetail, tempAddress}, ", " )

		receiveInfo.FullName = req.ReceiveInfo.FullName
		receiveInfo.Email = req.ReceiveInfo.Email
		receiveInfo.Phone = req.ReceiveInfo.Phone
		receiveInfo.Address = address
	} else {
		receiveInfo.FullName = userInfo.FullName
		receiveInfo.Email = userInfo.Email
		receiveInfo.Phone = userInfo.Phone
		receiveInfo.Address = userInfo.Address
	}
	now := utils.Now()
	bytes, _ := json.Marshal(receiveInfo)
	receiveInfoStr := string(bytes)
	// begin transaction
	var orderId uint
	err = models.Transaction(func(tx *gorm.DB) error {
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
			}
		}()
		//check coupon
		var couponId uint
		var exprired bool = false
		if req.Coupon != ""{
			couponId, exprired, err = checkCoupon(now, req.Coupon)
			if couponId <= 0{
				return errors.New("coupon is not valid")
			}
			if exprired {
				return errors.New("coupon is exprired")
			}
		}
		customerId := uint(userId)
		orderIns := models.Order{
			Status: mapStatusOrder[STATUS_PENDING],
			OrderDate: now,
			ReceiveInfo: &receiveInfoStr,
			CustomerID: &customerId,
			PaymentID: &req.PaymentMethod,
		}
		if err = tx.Create(&orderIns).Error; err != nil {
			return err
		}
		if orderIns.ID == 0{
			return errors.New("create order error")
		}
		orderId = orderIns.ID
		// create order detail
		for _, order := range req.Orders{
			orderDetail := models.OrderDetail{
				OrderID: orderId,
				BookID: uint(order.ProductID),
				Quantity: order.Amount,
				Price: order.Price,
			}
			if err = tx.Create(&orderDetail).Error; err != nil {
				return err
			}
		}

		if couponId > 0{
			//create order voucher
			orderVoucher := models.OrderVouchers{
				OrderID: orderId,
				VoucherID: couponId,
			}
			if err = tx.Create(&orderVoucher).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Info(err.Error())
		return responses.ErrSystem
	}
	order, _ := models.GetById(&models.Order{}, orderId)
	if order != nil {
		go events.BroadcastEvent(events.WsOrderCreated, order)
	}

	return nil
}

func checkCoupon(checkTime int64, code string) (id uint, expired bool, err error) {
	coupon, err := (&models.Voucher{}).GetByCode(code)
	if err != nil  {
		log.Info(err)
		if errors.Is(err, gorm.ErrRecordNotFound){
			return 0, false, errors.New("coupon is not existed")
		}
		return 0, false, responses.ErrSystem
	}

	return coupon.(*models.Voucher).ID, time.Unix(coupon.(*models.Voucher).Expiry, 0).After(time.Unix(checkTime, 0)), err
}

func OrderHistory(status, start, end, offset, limit int32) ([]data.OrderResponse, error) {
	var statusStr string = ""
	if val, ok  := mapStatusOrder[int(status)]; ok {
		statusStr = val
	}
	if end < 0 {
		end = 0
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
	orders, err := (&models.Order{}).GetWithFilter(statusStr, start, end, offset, limit)
	if err != nil {
		log.Info(err)
		return nil, responses.ErrSystem
	}
	return orders, nil
}