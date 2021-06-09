package utils
import (
	"os"
	"log"
)

func GetCurrentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		log.Println(err.Error(), "err.Error() services/userservice/user_service.go:563")
		return ""
	}

	return pwd
}