package controllers

import (
	"encoding/json"
	"go-tanxi/pkg/captcha"
	"go-tanxi/pkg/logger"
	"net/http"
)

type CaptchaController struct {
}

// Show 生成验证码
func (*CaptchaController) Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// 生成验证码
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogError(err)
	// 返回数据
	json.NewEncoder(w).Encode(map[string]interface{}{
		"code":          1,
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
