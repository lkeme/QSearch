// Package question 多选题
package question

import (
	"github.com/gin-gonic/gin"
	"github.com/lkeme/QSearch/app/controller/v1/baseapi"
	"github.com/lkeme/QSearch/app/response"
)

type BaseQuestionApi struct {
	baseapi.ResourceInterface // interface
}

func (bq *BaseQuestionApi) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (bq *BaseQuestionApi) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (bq *BaseQuestionApi) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (bq *BaseQuestionApi) Index(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (bq *BaseQuestionApi) Total(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (bq *BaseQuestionApi) Query(c *gin.Context) {
	keyword := c.Param("keyword")
	// err
	if keyword == "" && len(keyword) < 2 {
		response.FailWithMessage("关键字不能为空且不能小于两个长度", c)
		return
	}
}

func (bq *BaseQuestionApi) i() {}
