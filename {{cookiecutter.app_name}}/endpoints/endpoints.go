package endpoints

import (
	"net/http"
	"time"

	"%%baseimport%%/vars"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog/hlog"
)

// InitEndpoints 在全局变量初始化完成后进行
func InitEndpoints() error {
	rt := vars.Router

	// 安装一些中间件
	rt.Use(hlog.AccessHandler(func(r *http.Request, status, sz int, dur time.Duration) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Str("url", r.URL.String()).
			Int("status", status).
			Int("sz", size).
			Dur("dur", dur).
			Msg("")
	}))
	rt.Use(hlog.RequestIDHandler("reqid", "Inapp-Req-ID"))
	rt.Use(hlog.NewHandler(vars.Logger))
	rt.Use(middleware.Recoverer)

	rt.NotFound(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("找不到页面"))
	})

	return nil
}
