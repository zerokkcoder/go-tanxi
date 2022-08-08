package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.Header("Content-type", "text/html;charset=utf-8")
	c.String(http.StatusOK, "<h1>Helloooooo, 这里是 goblog</h1>")
}

func about(c *gin.Context) {
	c.Header("Content-type", "text/html;charset=utf-8")
	c.String(http.StatusOK, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func notFound(c *gin.Context) {
	c.Header("Content-type", "text/html;charset=utf-8")
	c.String(http.StatusNotFound, "<h1>请求页面未找到 :(</h1>"+
		"<p>如有疑惑，请联系我们。</p>")
}

func main() {
	// 1. 创建路由
	r := gin.New()
	// 2. 绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", index)
	r.GET("/about", about)
	r.NoRoute(notFound)
	// 3. 监听端口，默认在3000
	// Run("里面不指定端口号默认为3000")
	r.Run(":3000")
}
