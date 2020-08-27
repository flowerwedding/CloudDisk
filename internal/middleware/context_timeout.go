/**
 * @Title  context_timeout
 * @description  统一超时控制
 * @Author  沈来
 * @Update  2020/8/7 20:20
 **/
package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
