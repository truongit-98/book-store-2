
package responses


type ProductResponse struct {
	Id	uint  `json:"id"`
	Title string `json:"title"`
	Price	float64 `json:"price"`
	CoverImagge string `json:"cover_image"`
	Ratting float64 `json:"ratting"`
	Description   string  ` json:"description"`
	AuthorName string `json:"author_name"`
	TopicName	string `json:"topic_name"`
	Total	int		`json:"total"`
	TotalRow int `json:"total_row"`
	TotalReviewer int `json:"total_reviewer"`
}

type CategoryResponse struct {
	Id	uint  `json:"id"`
	CategoryName string `json:"category_name"`
	TopicId uint `json:"topic_id"`
	TopicName string `json:"topic_name"`
	TotalBook	int		`json:"total_book"`
}

type PublisherResponseWithCountBook struct {
	Id uint `json:"id"`
	PublisherName string `json:"publisher_name"`
	TotalBook int `json:"total_book"`
}
