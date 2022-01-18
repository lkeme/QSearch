//go:build windows
// +build windows

package bootstrap

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// windows
func initServer(address string, router *gin.Engine) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    2 * 60 * time.Second, // 10
		WriteTimeout:   2 * 60 * time.Second, // 10
		MaxHeaderBytes: 1 << 20,
	}
}
