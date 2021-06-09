package requestbody

type CategoryBody struct {
	CategoryName string `json:"category_name" xml:"category_name"`
}
type CategoryPutBody struct {
	ID uint `json:"id" xml:"id"`
	CategoryName string `json:"category_name" xml:"category_name"`
}
