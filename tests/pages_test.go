package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseURL = "http://localhost:3000"

// 请求 —— 模拟用户访问浏览器
var (
	resp *http.Response
	err  error
)

func TestHomePage(t *testing.T) {

	resp, err = http.Get(baseURL + "/")

	// 检测 —— 是否无错误且 200
	assert.NoError(t, err, "有错误发生，err 不为空")
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码 200")
}

func TestAboutPage(t *testing.T) {

	resp, err = http.Get(baseURL + "/about")

	// 检测 —— 是否无错误且 200
	assert.NoError(t, err, "有错误发生，err 不为空")
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码 200")
}
