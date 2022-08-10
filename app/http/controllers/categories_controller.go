package controllers

import (
	"fmt"
	"go-tanxi/app/models/category"
	"go-tanxi/app/requests"
	"go-tanxi/pkg/flash"
	"go-tanxi/pkg/route"
	"go-tanxi/pkg/view"
	"net/http"
)

// CategoriesController 文章分类控制器
type CategoriesController struct {
	BaseController
}

// Create 文章分类创建
func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

// Store 保存文章分类
func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	// 2. 表单验证
	errors := requests.ValidateCategoryForm(_category)

	// 3. 检查是否有错误
	if len(errors) == 0 {
		// 创建文章
		_category.Create()
		if _category.ID > 0 {
			flash.Success("分类创建成功")
			http.Redirect(w, r, route.Name2URL("articles.index"), http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")
	}
}

// Show 显示分类下的文章列表
func (*CategoriesController) Show(w http.ResponseWriter, r *http.Request) {

}
