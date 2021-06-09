package data

type OrderResponse struct {
	ID uint `json:"id"`
	Total float64 `json:"total"`
	DateCreated int64 `json:"date_created"`
	Status string `gorm:"size:50" json:"status"`
	OrderDate int64 `json:"order_date"`
	ReceiveInfo *string `json:"receive_info" orm:"models.ReceiveInfo"`
	CustomerID *uint `json:"customer_id"`
	PaymentID *uint `json:"payment_id"`
	Creator *uint `json:"creator"`
	Delivery *uint `json:"delivery"`
	TotalCount int `json:"total_count"`
}

