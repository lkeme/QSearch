package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/lkeme/QSearch/router"
	"github.com/thinkerou/favicon"
	"net/http"
	"strings"
)

func SetupRoute(r *gin.Engine) {
	// 配置静态资源
	setupStatic(r)

	// 注册全局中间件
	registerGlobalMiddleWare(r)

	//  注册 API 路由
	router.RegisterAPIRoutes(r)

	//  配置 404 路由
	//setup404Handler(r)
}

func registerGlobalMiddleWare(router *gin.Engine) {
	//router.Use(
	//middlewares.Logger(),
	//middlewares.Recovery(),
	//middlewares.ForceUA(),
	//)
}

func setupStatic(router *gin.Engine) {
	router.Use(favicon.New("./public/images/favicon.ico"))
	//router.Static("/assets", "./public/assets")
	//router.StaticFS("/static", http.Dir("./public/static"))
	//router.StaticFile("/favicon.ico", "./public/resources/favicon.ico")
}

func setup404Handler(router *gin.Engine) {
	// 处理 404 请求
	router.NoRoute(func(c *gin.Context) {
		// 获取标头信息的 Accept 信息
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			// 如果是 HTML 的话
			c.String(http.StatusNotFound, "页面返回 404")
		} else {
			// 默认返回 JSON
			c.JSON(http.StatusNotFound, gin.H{
				"error_code":    404,
				"error_message": "路由未定义，请确认 url 和请求方法是否正确。",
			})
		}
	})
}
