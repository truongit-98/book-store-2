package requestbody

type BookPostRequestBody struct {
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	CoverImage string `json:"cover_image"`
	Amount	int32	`json:"amount"`
	SKU	string `json:"sku"`
	TypeID uint `json:"book_type_id" `
}

type BookPutRequestBody struct {
	ID uint `json:"id"`
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	CoverImage string `json:"cover_image"`
	Price	float64	`json:"price"`
	Amount	int32	`json:"amount"`
	SKU	string `json:"sku"`
	BookTypeID uint `json:"book_type_id" `
}

type BookTypePostRequestBody struct{
	BookType int `json:"book_type"`
	TypeName string `json:"type_name"`
	Episodes int     `json:"episodes"`
}

type BookDetailPostRequestBody struct {
	PriceCover     float64 `json:"price_cover"`
	NumberOfPage   int    `json:"number_of_page"`
	Height         float64   `json:"height"`
	Width         float64 `json:"width"`
	Description    string  `gorm:"type:text" json:"description"`
	Language       string  `gorm:"size:20" json:"language"`
	BookTypeID         uint `json:"book_type_id"`
	PublisherID    uint    `json:"publisher_id"`
	FormatID       uint    `json:"format_id"`
	TopicID uint            `json:"topic_id"`
}