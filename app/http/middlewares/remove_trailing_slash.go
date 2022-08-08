package middlewares

import (
	"net/http"
	"strings"
)

// 去除 uri 中最后一个斜杠 /
func RemoveTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. 除首页以外，移除所有请求路径后面的斜杆
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// 2. 继续处理请求
		next.ServeHTTP(w, r)
	})
}
