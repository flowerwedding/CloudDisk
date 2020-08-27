/**
 * @Title  captcha
 * @description  #
 * @Author  沈来
 * @Update  2020/8/24 9:15
 **/
package api

import (
	"CloudDisk/pkg/captcha/model"
	"CloudDisk/pkg/captcha/recaptcha"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

// @Summary  生成图形验证码
// @Produce  json
// @Success  200 {object} string "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /getCaptcha [get]
func GetCaptcha(context *gin.Context) {
	baseResponse := model.NewBaseResponse()
	d := struct {
		CaptchaId string
	}{
		captcha.New(),
	}
	if d.CaptchaId != "" {
		baseResponse.GetSuccessResponse()
		var captcha model.CaptchaResponse
		captcha.CaptchaId = d.CaptchaId
		captcha.ImageUrl = "/show/" + d.CaptchaId + ".png"
		baseResponse.Data = captcha
	} else {
		baseResponse.GetFailureResponse(model.SYSTEM_ERROE)
	}
	context.JSON(http.StatusOK, baseResponse)
}

// @Summary  验证图形验证码
// @Produce  json
// @Param captchaId body string true "验证码ID"
// @Param value body string true "验证码内容"
// @Success  200 {object} string "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /auth [get]
func VerifyCaptcha(context *gin.Context) {
	baseResponse := model.NewBaseResponse()
	captchaId := context.Request.FormValue("captchaId")
	value := context.Request.FormValue("value")
	if captchaId == "" || value == "" {
		baseResponse.GetFailureResponse(model.QUERY_PARAM_ERROR)
	} else {
		if captcha.VerifyString(captchaId, value) {
			baseResponse.GetSuccessResponse()
			baseResponse.Message = "验证成功"
		} else {
			baseResponse.GetFailureResponse(model.CAPTCHA_ERROR)
		}
	}
	context.JSON(http.StatusOK, baseResponse)
}

// @Summary  查看图形验证码
// @Produce  json
// @Success  200 {object} string "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /show/:source [get]
func GetCaptchaPng(context *gin.Context) {
	source := context.Param("source")
	logrus.Info("GetCaptchaPng : " + source)
	recaptcha.ServeHTTP(context.Writer, context.Request)
}
