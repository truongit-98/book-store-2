package requestbody

type WSBodyRequest struct {
	Channel string `json:"channel" xml:"channel"`
	Message string `json:"message" xml:"message"`
}

