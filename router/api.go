package router

import (
	"github.com/gin-gonic/gin"
	api "github.com/lkeme/QSearch/app/controller/v1"
	"github.com/lkeme/QSearch/app/middleware"
	_ "github.com/lkeme/QSearch/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// Default() middleware New() No middleware

	// SwagAPI register swagger handler
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// unified path
	v1 := r.Group("/api/v1")

	// No authentication required
	publicGroup := v1.Group("")
	{
		// SystemRouterGroup
		systemGroup := publicGroup.Group("system")
		{
			// /api/v1/system/health
			systemGroup.GET("/health", api.ApiGroupApp.SystemApiGroup.Health)
		}

	}

	// Authentication required
	privateGroup := v1.Group("")
	privateGroup.Use(middleware.JWTAuth())
	{

	}
}
