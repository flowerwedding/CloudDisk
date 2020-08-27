/**
 * @Title  jwt
 * @description  Token中间件
 * @Author  沈来
 * @Update  2020/8/7 14:48
 **/
package middleware

import (
	"CloudDisk/pkg/app"
	"CloudDisk/pkg/errcode"
	"CloudDisk/pkg/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)

		//postman的header里面加token {{Token}}
		if s, exist := c.GetQuery("token"); exist {
			token = s
		} else {
			token = c.GetHeader("token")
		}
		if token == "" {
			ecode = errcode.InvalidParams
		} else {
			claim, err := app.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
			if claim != nil {
				c.Set("user", util.Decode(claim.Username))
			}
		}

		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
		}

		c.Next()
	}
}
