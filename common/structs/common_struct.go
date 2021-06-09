package structs

type AccessDetails struct {
	AccessUuid string
	User string
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type BookDataFake struct {
	BookID int `csv:"bookID"`
	Title string `csv:"title"`
	Authors string  `csv:"authors"`
	Average_rating string `csv:"average_rating"`
	Isbn string `csv:"isbn"`
	Isbn13 string `csv:"isbn13"`
	Language_code string `csv:"language_code"`
	Num_pages string  `csv:"num_pages"`
	Ratings_count string `csv:"ratings_count"`
	Text_reviews_count string `csv:"text_reviews_count"`
	Publication_date string  `csv:"publication_dateID"`
	Publisher string `csv:"publisher"`
} 