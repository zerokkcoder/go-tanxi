package controllers

import (
	"go-tanxi/pkg/view"
	"net/http"
)

// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "pages.home")
}

// About 关于
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "pages.about")
}

// NotFound 404 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	view.RenderSimple(w, view.D{}, "pages.not_found")
}
