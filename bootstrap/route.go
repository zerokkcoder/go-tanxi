package bootstrap

import (
	"embed"
	"go-tanxi/pkg/route"
	"go-tanxi/routes"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute(staticFS embed.FS) *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)

	route.SetRoute(router)

	// 静态资源
	sub, _ := fs.Sub(staticFS, "public")
	router.PathPrefix("/").Handler(http.FileServer(http.FS(sub)))

	return router
}
