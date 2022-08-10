package middlewares

import (
	"go-tanxi/pkg/auth"
	"go-tanxi/pkg/flash"
	"go-tanxi/pkg/route"
	"net/http"
)

// Auth 登录用户才可访问
func Auth(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !auth.Check() {
			flash.Warning("登录用户才能访问此页面")
			http.Redirect(w, r, route.Name2URL("articles.index"), http.StatusFound)
			return
		}

		next(w, r)
	}
}
