package requestbody


type order struct {
	ProductID int `json:"product_id"`
	Amount int `json:"amount"`
	Price float64 `json:"price"`
}

type receiveInfo struct {
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	CityId string `json:"city_id"`
	DistrictId string `json:"district_id"`
	WardId string `json:"ward_id"`
	AddressDetail string `json:"address_detail"`
}

type OrderInformation struct {
	ReceiveInfo *receiveInfo `json:"receive_info"`
	Orders []*order `json:"orders"`
	PaymentMethod uint `json:"payment_method"`
	Coupon string `json:"coupon"`
}