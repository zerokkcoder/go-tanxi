package controllers

import (
	"encoding/json"
	"fmt"
	"go-tanxi/app/models/user"
	"go-tanxi/app/requests"
	"go-tanxi/pkg/route"
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

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 2. 表单规则
	errs := requests.ValidateRegistrationForm(_user)

	if len(errs) > 0 {
		// 4.1 有错误发生，打印数据
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
		data, _ := json.MarshalIndent(errs, "", "  ")
		fmt.Fprint(w, string(data))
	} else {
		_user.Create()
		if _user.ID > 0 {
			http.Redirect(w, r, route.Name2URL("articles.index"), http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建用户失败，请联系管理员")
		}
	}
}
