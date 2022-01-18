package main

import (
	"github.com/lkeme/QSearch/bootstrap"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w export GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
//go:generate go mod tidy
//go:generate go mod download

func main() {
	// RunServer
	bootstrap.RunServer()
}
