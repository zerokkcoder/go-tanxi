package middlewares

import (
	"go-tanxi/pkg/auth"
	"go-tanxi/pkg/flash"
	"go-tanxi/pkg/route"
	"net/http"
)

// Guest 只允许未登录用户访问
func Guest(next HttpHandlerFunc) HttpHandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if auth.Check() {
			flash.Warning("登录用户无法访问此页面")
			http.Redirect(w, r, route.Name2URL("articles.index"), http.StatusFound)
			return
		}

		next(w, r)
	}
}
