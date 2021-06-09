package utils

import (
	"BookStore/common/structs"
	"BookStore/consts"
	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
	"github.com/twinj/uuid"
	"time"
	"fmt"
)

const (
	ACCOUNT_ADMIN = "ACCOUNT_ADMIN"
	ACCOUNT_USER = "ACCOUNT_USER"
)


func CreateToken(accountId int, typeAccount string) (*structs.TokenDetails, error) {
	var err error
	//Creating Access Token
	token := &structs.TokenDetails{
		AtExpires: time.Now().Add(time.Hour * 24 * 7).Unix(),
		AccessUuid: uuid.NewV4().String(),
		RtExpires: time.Now().Add(time.Hour * 24 * 7).Unix(),
		RefreshUuid: uuid.NewV4().String(),
	}
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = token.AccessUuid
	atClaims["user"] = fmt.Sprintf("%d", accountId)
	atClaims["type_account"] = typeAccount
	atClaims["exp"] = token.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token.AccessToken, err = at.SignedString([]byte(consts.ACCESS_SECRET))
	if err != nil {
		log.Info(err.Error())
	   return nil, err
	}
	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = token.RefreshUuid
	rtClaims["user"] = fmt.Sprintf("%d", accountId)
	atClaims["type_account"] = typeAccount
	rtClaims["exp"] = token.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	token.RefreshToken, err = rt.SignedString([]byte(consts.REFRESH_SECRET))
	if err != nil {
		log.Info(err.Error())
		return nil, err
	}
	return token, nil
  }

