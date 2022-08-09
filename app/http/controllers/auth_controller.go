package controllers

import (
	"go-tanxi/pkg/view"
	"net/http"
)

// AuthController 用户认证
type AuthController struct {
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 注册
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {

}
