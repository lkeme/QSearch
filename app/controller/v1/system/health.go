package system

import (
	"github.com/gin-gonic/gin"
	"github.com/lkeme/QSearch/app/controller/v1/baseapi"
	"github.com/lkeme/QSearch/app/response"
)

type HealthSystem struct {
	baseapi.BaseController
}

func (h *HealthSystem) Health(c *gin.Context) {
	response.Success(c)
}
