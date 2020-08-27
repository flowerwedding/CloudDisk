/**
 * @Title  access_log
 * @description  访问日志记录
 * @Author  沈来
 * @Update  2020/8/7 16:41
 **/
package middleware

import (
	"CloudDisk/global"
	"CloudDisk/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Writer(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: c.Writer,
		}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		global.Logger.WithFields(fields).Infof(c, "access log: method: %s, status_code: %d, begin_time: %d, end_time: %d",
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
