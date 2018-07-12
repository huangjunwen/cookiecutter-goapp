package endpoints

import (
	"net/http"

	"%%baseimport%%/vars"
	"github.com/huangjunwen/cookiecutter-goapp/utils/zlog"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
)

// InitEndpoints 在全局变量初始化完成后进行
func InitEndpoints() error {
	rt := vars.Router

	// 一些基础中间件，从上到下对应从最外层到最内层
	rt.Use(hlog.NewHandler(vars.Logger))
	rt.Use(hlog.RequestIDHandler("reqid", "Inapp-Req-ID"))
	rt.Use(middleware.RequestLogger(zlog.ZLogFormatter{}))
	rt.Use(middleware.Recoverer)

	rt.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello {{cookiecutter.app_name}}"))
	})

	return nil
}
