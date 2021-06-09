package redisservice

import (
	"BookStore/common/structs"
	"github.com/go-redis/redis"
	"github.com/prometheus/common/log"
	"time"
	"fmt"
)

var  client *redis.Client

func init() {
	//Initializing redis
	dsn := "aaa"
	//if len(dsn) == 0 {
	dsn = "localhost:6379"
	//}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func CreateAuth(user uint, token *structs.TokenDetails) error {
	at := time.Unix(token.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(token.RtExpires, 0)
	now := time.Now()

	errAccess := client.Set(token.AccessUuid, fmt.Sprintf("%d", user), at.Sub(now)).Err()
	if errAccess != nil {
		log.Info(errAccess.Error())
		return errAccess
	}
	errRefresh := client.Set(token.RefreshUuid, fmt.Sprintf("%d", user), rt.Sub(now)).Err()
	if errRefresh != nil {
		log.Info(errRefresh.Error())
		return errRefresh
	}
	return nil
}

func FetchAuth(accessUuid string) (*string, error) {
	user, err := client.Get(accessUuid).Result()
	if err != nil {
		return nil, err
	}
	return &user, nil
}