/**
 * @Title  enforcer
 * @description  拦截器
 * @Author  沈来
 * @Update  2020/8/25 21:37
 **/
package middleware

import (
	"CloudDisk/pkg/app"
	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
	"strings"
)

func Authorize(sub string, e *casbin.Enforcer) gin.HandlerFunc {

	return func(c *gin.Context) {
		url := strings.Split(c.Request.URL.RequestURI(), "?")
		obj := url[0] + "/" + c.Query("username")
		act := c.Request.Method
		response := app.NewResponse(c)
		var message string
		if ok := e.Enforce(sub, obj, act); ok {
			message = "恭喜您,权限验证通过"
			c.Next()
		} else {
			message = "很遗憾,权限验证没有通过"
			c.Abort()
		}

		response.ToResponse(gin.H{
			"message": message,
		})
	}
}
