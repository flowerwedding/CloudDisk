/**
 * @Title  app_info
 * @description  服务信息存储
 * @Author  沈来
 * @Update  2020/8/7 18:50
 **/
package middleware

import "github.com/gin-gonic/gin"

func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "myBlog")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
