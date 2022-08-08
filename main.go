package main

import (
	"go-tanxi/app/http/middlewares"
	"go-tanxi/bootstrap"
	"go-tanxi/pkg/logger"
	"net/http"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
