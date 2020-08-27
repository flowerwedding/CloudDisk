/**
 * @Title  recovery
 * @description  自定义recovery
 * @Author  沈来
 * @Update  2020/8/7 17:03
 **/
package middleware

import (
	"CloudDisk/global"
	"CloudDisk/pkg/app"
	"CloudDisk/pkg/email"
	"CloudDisk/pkg/errcode"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailSetting.Host,
		Port:     global.EmailSetting.Port,
		IsSSL:    global.EmailSetting.IsSSL,
		UserName: global.EmailSetting.UserName,
		Password: global.EmailSetting.Password,
		From:     global.EmailSetting.From,
	})

	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf("panic recover err: %v", err)

				err := defaultMailer.SendEmail(
					global.EmailSetting.To,
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息：%v", err),
				)
				if err != nil {
					global.Logger.Panicf("mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
