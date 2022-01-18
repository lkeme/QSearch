//go:build !windows
// +build !windows

package bootstrap

import (
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

// non windows
func initServer(address string, router *gin.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
