package requestbody


type UserLoginBody struct {
	Email string ` json:"email"`
	Password string ` json:"password"`
}

type UserRequestBody struct {
	Id uint `json:"id"`
	Password string ` json:"password"`
	Email string ` json:"email"`
	FullName string `json:"full_name"`
	Address string ` json:"address"`
	Phone string ` json:"phone"`
}