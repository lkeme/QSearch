package bootstrap

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lkeme/QSearch/pkg/config"
)

type server interface {
	ListenAndServe() error
}

var route *gin.Engine

func RunServer() {
	// 初始化server
	address := ":" + config.Get("app.port")
	s := initServer(address, route)
	// 欢迎词
	fmt.Printf(`
		欢迎使用: %s - %s
		默认启动地址地址: %s
	`, config.Get("app.name"), config.Get("app.version"), config.Get("app.url"),
	)
	// 运行
	s.ListenAndServe().Error()
}

func init() {
	// support: debug, release, test
	gin.SetMode(gin.ReleaseMode)
	// config
	config.InitConfig()
	// route
	route = gin.Default()
	// router
	SetupRoute(route)
	// database
	SetupDatabase()

}
