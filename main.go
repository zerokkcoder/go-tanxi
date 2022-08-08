package main

import (
	"database/sql"
	"fmt"
	"go-tanxi/bootstrap"
	"go-tanxi/pkg/database"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// 1. 创建路由
var router = mux.NewRouter()
var db *sql.DB

// 中间件 设置 Content-Type: text/html; charset=utf-8
func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 设置标头
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

// 去除 uri 中最后一个斜杠 /
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}

func main() {
	database.Initialize()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	// 中间件 设置 Content-Type: text/html; charset=utf-8
	router.Use(forceHTMLMiddleware)

	// 通过命名路由获取 URL 示例
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "1")
	fmt.Println("articleURL: ", articleURL)

	// 3. 监听端口，默认在3000
	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
