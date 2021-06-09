package middleware

import (
	"BookStore/common/structs"
	"BookStore/consts"
	"BookStore/restapi/responses"
	"BookStore/services/adminservice"
	"BookStore/services/responseservice"
	"BookStore/utils"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"
	"github.com/prometheus/common/log"
	"net/http"
	"strconv"
	"strings"
)
//func CheckSession() beego.FilterFunc {
//	return func(ctx *context.Context) {
//		if !strings.Contains(ctx.Request.URL.String(), "/v1/api/account/admins/") && !strings.Contains(ctx.Request.URL.String(), "/v1/api/account/admins/login") && !strings.Contains(ctx.Request.URL.String(), "/v1/api/account/admins/register") {
//			tokenAuth, err := ExtractTokenMetadata(ctx.Request)
//			if err != nil {
//				log.Info(err.Error())
//				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.ErrUnknown), true, true)
//				ctx.ResponseWriter.WriteHeader(http.StatusOK)
//				return
//			}
//			user, err := redisservice.FetchAuth(tokenAuth.AccessUuid)
//			if err != nil {
//				log.Info(err.Error())
//				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.UnAuthorized), true, true)
//				ctx.ResponseWriter.WriteHeader(http.StatusOK)
//				return
//			}
//			if *user != tokenAuth.User {
//				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.UnAuthorized), true, true)
//				ctx.ResponseWriter.WriteHeader(http.StatusOK)
//				return
//			}
//			ctx.Input.SetParam("user", *user)
//			return
//		}
//		return
//	}
//}

var mapNotAuthor = map[string]string{
	"/v1/api/wsorders": "",
	"/v1/api/order/checkout": "",
	"/v1/api/vouchers/": "",
	"/v1/api/account/admins/login": "",
	"/v1/api/account/customer/login": "",

	"/v1/api/books/featured": "",
	"/v1/api/books/filter": "",
	"/v1/api/books/info": "",
	"/v1/api/books/new": "",
	"/v1/api/books/seller": "",
	"/v1/api/common": "",
	"/v1/api/payments": "",
	"/v1/api/publishers/with-book-count": "",
	"/v1/api/categories/with-book-count": "",

}

func checkUri(uri string) bool {
	for u, _ := range mapNotAuthor {
		if strings.Contains(uri, u){
			return true
		}
	}
	return false
}

func CheckPermission() beego.FilterFunc {
	return func(ctx *context.Context) {
		var path = ctx.Request.URL.Path
		var action = ctx.Request.Method
		if ok := checkUri(path); ok {
			return
		} else {
			tokenAuth, typeAccount, err := ExtractTokenMetadata(ctx.Request)
			if err != nil {
				log.Info(err.Error())
				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.BadRequest), true, true)
				ctx.ResponseWriter.WriteHeader(http.StatusOK)
				return
			}
			if typeAccount == utils.ACCOUNT_USER{
				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.UnAuthorized), true, true)
				ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
				return
			}
			accountId, err := strconv.Atoi(tokenAuth.User)
			if err != nil {
				log.Info(err.Error())
				ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.BadRequest), true, true)
				ctx.ResponseWriter.WriteHeader(http.StatusOK)
				return
			}

			authorInfo, err := adminservice.GetAuthorInfoForAdmin(int32(accountId))
			if err != nil {
				log.Info(err.Error())
				ctx.Output.JSON(responseservice.GetCommonErrorResponse(err), true, true)
				ctx.ResponseWriter.WriteHeader(http.StatusOK)
				return
			}
			ctx.Input.SetParam("adminId", strconv.Itoa(accountId))
			if accountId == 1 {
				return
			}
			if strings.HasPrefix(path, "/v1/api/account/admins/author-info"){
				return
			}
			//authors, err := (&models.Role{}).GetAllWithPreload()
			//if err != nil {
			//	log.Info(err.Error())
			//	ctx.Output.JSON(responseservice.GetCommonErrorResponse(err), true, true)
			//	ctx.ResponseWriter.WriteHeader(http.StatusOK)
			//	return
			//}
			for _, r := range authorInfo.Roles{
				for _, p := range r.Permissions{
					if strings.HasPrefix(path, p.Path){
						for _, c := range p.Actions{
							if strings.ToUpper(c.Action) == strings.ToUpper(action){
								return
							}
						}
					}
				}
			}
			ctx.Output.JSON(responseservice.GetCommonErrorResponse(responses.UnAuthorized), true, true)
			ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
			return
			log.Info(authorInfo)
		}

		return
	}
}

func ExtractTokenMetadata(r *http.Request) (*structs.AccessDetails, string, error) {
	token, err := verifyToken(r)
	if err != nil {
		return nil, "", errors.New("token is not valid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, "", errors.New("token is not valid")
		}
		user, ok := claims["user"].(string)
		if !ok {
			return nil, "", errors.New("token is not valid")
		}
		typeAccount, ok := claims["type_account"].(string)
		if !ok {
			return nil, "", errors.New("token is not valid")
		}
		return &structs.AccessDetails{
			AccessUuid: accessUuid,
			User:   user,
		}, typeAccount, nil
	}
	return nil, "", err
}
func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := r.Header.Get("token")
	log.Info(tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(consts.ACCESS_SECRET), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
