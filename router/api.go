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

	// unified path|check sql inject
	v1 := r.Group("/api/v1", middleware.CheckSQLInject())

	// No authentication required
	publicGroup := v1.Group("")
	{
		// SystemRouterGroup
		systemGroup := publicGroup.Group("system")
		{
			// /api/v1/system/health
			systemGroup.GET("/health", api.ApiGroupApp.SystemApiGroup.Health)
		}

		// Question System
		questionGroup := publicGroup.Group("question")
		{
			questionGroup.POST("", api.ApiGroupApp.QuestionApiGroup.BaseQuestionApi.Create)
			questionGroup.DELETE("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionApi.Delete)
			questionGroup.PUT("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionApi.Update)
			questionGroup.GET("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionApi.Index)
			questionGroup.GET("/total", api.ApiGroupApp.QuestionApiGroup.BaseQuestionApi.Total)
			questionGroup.GET("/:keyword", api.ApiGroupApp.QuestionApiGroup.BaseQuestionApi.Query)
		}
		// Question Tag System
		questionTagGroup := publicGroup.Group("question")
		{
			questionTagGroup.POST("", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTagApi.Create)
			questionTagGroup.DELETE("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTagApi.Delete)
			questionTagGroup.PUT("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTagApi.Update)
			questionTagGroup.GET("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTagApi.Index)
			questionTagGroup.GET("/total", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTagApi.Total)
			questionTagGroup.GET("/:keyword", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTagApi.Query)
		}
		// Question Type System
		questionTypeGroup := publicGroup.Group("question")
		{
			questionTypeGroup.POST("", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTypeApi.Create)
			questionTypeGroup.DELETE("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTypeApi.Delete)
			questionTypeGroup.PUT("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTypeApi.Update)
			questionTypeGroup.GET("/:id", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTypeApi.Index)
			questionTypeGroup.GET("/total", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTypeApi.Total)
			questionTypeGroup.GET("/:keyword", api.ApiGroupApp.QuestionApiGroup.BaseQuestionTypeApi.Query)
		}

	}

	// Authentication required
	privateGroup := v1.Group("")
	privateGroup.Use(middleware.JWTAuth())
	{

	}
}
