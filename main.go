package main

import (
	"database/sql"
	"go-tanxi/app/http/middlewares"
	"go-tanxi/bootstrap"
	"go-tanxi/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

// 1. 创建路由
var router = mux.NewRouter()
var db *sql.DB

func main() {
	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
