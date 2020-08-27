/**
 * @Title  jwt
 * @description  鉴权
 * @Author  沈来
 * @Update  2020/8/6 21:44
 **/
package app

import (
	"CloudDisk/global"
	"CloudDisk/pkg/util"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(global.JWTSetting.Expire)
	claims := Claims{
		Username: util.Encode(username),
		Password: util.Encode(password),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    global.JWTSetting.Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//这个真的好坑..，SignedString里面参数是[]byte，要查jwt方面查不能从报错查
	token, err := tokenClaims.SignedString(GetJWTSecret())
	fmt.Println(token, err)
	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return GetJWTSecret(), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
