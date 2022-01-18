package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Unauthorized 响应 401，未传参 msg 时使用默认消息
// 登录失败、jwt 解析失败时调用
func Unauthorized(c *gin.Context, msg ...string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。", msg...),
	})
}

// defaultMessage 内用的辅助函数，用以支持默认参数默认值
// Go 不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string) (message string) {
	if len(msg) > 0 {
		message = msg[0]
	} else {
		message = defaultMsg
	}
	return
}
