/**
 * @Title  auth
 * @description  JWT
 * @Author  沈来
 * @Update  2020/8/6 22:13
 **/
package api

import (
	"CloudDisk/global"
	"CloudDisk/internal/service"
	"CloudDisk/pkg/app"
	"CloudDisk/pkg/errcode"
	"github.com/gin-gonic/gin"
)

// @Summary  登录获取JWT
// @Produce  json
// @Param  username query string true "Key"
// @Param  password query string true "Secret"
// @Success  200 {object} model.Auth "成功"
// @Failure  400 {object} errcode.Error "请求错误"
// @Failure  500 {object} errcode.Error "内部错误"
// @Router  /auth [get]
func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CheckAuth(&param)
	if err != nil {
		global.Logger.Errorf("svc.CheckAuth err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.Username, param.Password)
	if err != nil {
		global.Logger.Errorf("app.GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token":   token,
		"message": "登录成功",
	})
}
