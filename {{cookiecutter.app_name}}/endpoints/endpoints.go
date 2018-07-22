package endpoints

import (
	"net/http"

	"%%baseimport%%/vars"
	"bitbucket.org/jayven/platform-kit/utils/zlog"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

// InitEndpoints 在全局变量初始化完成后进行
func InitEndpoints() error {
	rt := vars.Router

	// 一些基础中间件，从上到下对应从最外层到最内层
	rt.Use(hlog.NewHandler(vars.Logger))
	rt.Use(middleware.RequestID)
	rt.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			reqID := middleware.GetReqID(ctx)
			if reqID != "" {
				log := zerolog.Ctx(ctx)
				log.UpdateContext(func(c zerolog.Context) zerolog.Context {
					return c.Str(vars.ReqIDFieldName, reqID)
				})
			}
			next.ServeHTTP(w, r)
		})
	})
	rt.Use(middleware.RequestLogger(zlog.ZLogFormatter{}))
	rt.Use(middleware.Recoverer)

	rt.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello {{cookiecutter.app_name}}"))
	})

	return nil
}
