package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lkeme/QSearch/app/middleware"
)

func RegisterAPIRoutes(r *gin.Engine) {
	// Default() middleware New() No middleware

	// unified path
	v1 := r.Group("/api/v1")

	// No authentication required
	publicGroup := v1.Group("")
	{
		// /api/v1/health  health check
		publicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})

	}

	// Authentication required
	privateGroup := v1.Group("")
	privateGroup.Use(middleware.JWTAuth())
	{

	}
}
