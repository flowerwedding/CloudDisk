/**
 * @Title  enforcer
 * @description  #
 * @Author  沈来
 * @Update  2020/8/25 20:26
 **/
package api

import (
	"CloudDisk/global"
	"CloudDisk/pkg/app"
	"fmt"
	"github.com/gin-gonic/gin"
)

func EnforcerAdd(c *gin.Context) {
	message := "增加Policy,"
	response := app.NewResponse(c)
	obj := "/auth/" + c.Query("admin")

	if ok := global.Enforcer.AddPolicy("admin", obj, "GET"); !ok {
		message += "Policy已经存在"
	} else {
		message += "增加成功"
	}

	response.ToResponse(gin.H{
		"message": message,
	})
}

func EnforcerDelete(c *gin.Context) {
	message := "删除Policy,"
	response := app.NewResponse(c)
	obj := "/auth/" + c.Query("admin")

	fmt.Println(obj)
	if ok := global.Enforcer.RemovePolicy("admin", obj, "GET"); !ok {
		message += "Policy不存在"
	} else {
		message += "删除成功"
	}

	response.ToResponse(gin.H{
		"message": message,
	})
}

func EnforcerGet(c *gin.Context) {
	var message []string
	message = append(message, "查看policy,")
	response := app.NewResponse(c)

	list := global.Enforcer.GetPolicy()
	for _, vlist := range list {
		var vv string
		for _, v := range vlist {
			vv += v + " "
		}
		message = append(message, fmt.Sprintf("value: %s, ", vv))
	}

	response.ToResponse(gin.H{
		"message": message,
	})
}
