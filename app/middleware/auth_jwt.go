package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lkeme/QSearch/app/model/user"
	"github.com/lkeme/QSearch/app/response"
	"github.com/lkeme/QSearch/pkg/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// parse
		claims, err := jwt.NewJWT().ParserToken(c)
		// error
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("401 Unauthorized: %v", err.Error()))
		}
		// ok
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}
		// set gin.context
		c.Set("claims", claims)
		c.Next()

	}
}
