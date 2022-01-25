package baseapi

import "github.com/gin-gonic/gin"

type BaseController struct{}

// 实现编译期间检测接口是否实现
var _ ResourceInterface = (*ResourceStruct)(nil)

// ResourceStruct Base struct
type ResourceStruct struct{}

// ResourceInterface Base interface
type ResourceInterface interface {
	i()                    //
	Create(c *gin.Context) //
	Delete(c *gin.Context) //
	Update(c *gin.Context) //
	Index(c *gin.Context)  //
	Total(c *gin.Context)  //
	Query(c *gin.Context)  //
}

// i 为了避免被其他包实现
func (r *ResourceStruct) i() {}

func (r *ResourceStruct) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *ResourceStruct) Delete(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *ResourceStruct) Update(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *ResourceStruct) Index(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *ResourceStruct) Total(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (r *ResourceStruct) Query(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}
