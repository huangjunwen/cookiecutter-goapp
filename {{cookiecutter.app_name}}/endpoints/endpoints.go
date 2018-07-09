package endpoints

import (
	"net/http"

	"%%baseimport%%/vars"
)

// InitEndpoints 在全局变量初始化完成后进行
func InitEndpoints() error {
	r := vars.Router

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("找不到页面"))
	})

	return nil
}
