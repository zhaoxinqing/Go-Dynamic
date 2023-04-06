package backend

import (
	"backend-go/app/backend/handler"
	"backend-go/public"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(addr string) {

	gin.SetMode(gin.ReleaseMode) // DebugMode、ReleaseMode、TestMode

	engine := gin.Default()

	engine.SetTrustedProxies(nil) // 代理

	routers(engine) // 路由注册

	// time.Sleep(10 * time.Second)

	engine.Run(addr)
}

func routers(app *gin.Engine) {

	api := app.Group("web-api")

	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"msg": time.Now().Format(public.TIME_FORMAT_ZONE)})
	})

	api.GET("some-json", handler.SomeJson)
	// 基于社会实时资讯，调整优化自己生活的旋律，丰富自己的财富及资本，有固守不变的核心，也有适应性调整的更新。变的是外在的形式，不变的是核心；
}
