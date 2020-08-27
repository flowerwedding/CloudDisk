/**
 * @Title  translations
 * @description  语言包翻译
 * @Author  沈来
 * @Update  2020/8/4 15:39
 **/
package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

func Translations() gin.HandlerFunc{
	return func(c *gin.Context){
		//这里路径重复、缺失方法是因为vendor里面没有，重新go mod拉包就行了
		uni := ut.New(en.New(),zh.New(), zh_Hant_TW.New())
		local := c.GetHeader("local")
		trans, _ := uni.GetTranslator(local)
		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch local{
			case "zh":
				_ = zh_translations.RegisterDefaultTranslations(v,trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v,trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v,trans)
				break
			}
			c.Set("trans",trans)
		}
		c.Next()
	}
}