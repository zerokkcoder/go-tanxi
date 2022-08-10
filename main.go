package main

import (
	"go-tanxi/app/http/middlewares"
	"go-tanxi/bootstrap"
	"go-tanxi/config"
	pkgConfig "go-tanxi/pkg/config"
	"go-tanxi/pkg/logger"
	"net/http"
)

func init() {
	// 初始化配置信息
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":"+pkgConfig.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
