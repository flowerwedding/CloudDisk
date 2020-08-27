/**
 * @Title  router
 * @Description  路由管理
 * @Author  沈来
 * @Update  2020/8/3 16:05
 **/
package routers

import (
	//	_ "CloudDisk/docs"
	"CloudDisk/global"
	"CloudDisk/internal/middleware"
	"CloudDisk/internal/routers/api"
	v1 "CloudDisk/internal/routers/api/v1"
	"net/http"

	//	v1 "CloudDisk/internal/routers/api/v1"
	"CloudDisk/pkg/limiter"
	"github.com/gin-gonic/gin"
	//	ginSwagger "github.com/swaggo/gin-swagger"
	//	"github.com/swaggo/gin-swagger/swaggerFiles"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()

	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())

	//r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ap := r.Group("/api") //默认管理员username = hello
	ap.Use(middleware.Authorize("charge", global.Enforcer))
	{
		ap.POST("/add", api.EnforcerAdd)
		ap.DELETE("/delete", api.EnforcerDelete)
		ap.GET("/get", api.EnforcerGet)
	}

	r.GET("/getCaptcha", api.GetCaptcha) //图形验证码
	r.GET("/verifyCaptcha", api.VerifyCaptcha)
	r.GET("/show/:source", api.GetCaptchaPng)

	//权限合法才能登录成功，登录成功才能调用下面带token的操作
	r.GET("/auth", middleware.Authorize("admin", global.Enforcer), api.GetAuth)

	file := v1.NewFile()

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		//任何请求输入文件名加后缀
		apiv1.POST("/file", file.Create)
		apiv1.GET("/file/link", file.GetByLink)
		apiv1.GET("/file/Qt", file.GetByQt)
		apiv1.GET("/file", file.Download)
		r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath)) //必须有,static URL才能被访问
	}

	return r
}
