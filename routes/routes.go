package routes

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middlewares"
	"bluebell/setting"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // gin设置成发布模式
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册
	r.POST("/signup", controller.SignUpHandler)
	// 登录
	r.POST("/login", controller.LoginHandler)

	r.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, setting.Conf.Version)
	})

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录的用户,判断请求头中是否有 有效的JWT  ？
		c.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
