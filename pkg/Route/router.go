package route

import (
	"go-tanxi/pkg/logger"
	"go-tanxi/routes"
	"net/http"

	"github.com/gorilla/mux"
)

// Route 路由对象
var Router *mux.Router

// Initialize 初始化路由
func Initialize() {
	Router = mux.NewRouter()
	// 注册路由
	routes.RegisterWebRoutes(Router)
}

// RouteName2URL 通过路由名称来获取 URL
func Name2URL(routeName string, pairs ...string) string {
	url, err := Router.Get(routeName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	}

	return url.String()
}

// GetRouteVariable 获取 URI 路由参数
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}
