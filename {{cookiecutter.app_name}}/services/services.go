package services

import (
	"context"

	"%%baseimport%%/vars"
	"github.com/go-chi/chi/middleware"
	"github.com/huangjunwen/platform-kit/svc"
	"github.com/rs/zerolog"
)

// InitServices 在全局变量初始化完成后进行
func InitServices() error {
	// 给 SvcServer/SvcClient 安装中间件
	vars.SvcServer = libsvc.DecorateServer(vars.SvcServer, func(h libsvc.ServiceHandler) libsvc.ServiceHandler {
		return func(ctx context.Context, method libsvc.Method, input, output interface{}) error {
			// 创建一个 logger 的拷贝
			l := vars.Logger.With().Logger()

			// 是否有传递过来的 reqid
			reqID := libsvc.Passthru(ctx)[vars.ReqIDFieldName]
			if reqID != "" {
				l.UpdateContext(func(c zerolog.Context) zerolog.Context {
					return c.Str(vars.ReqIDFieldName, reqID)
				})
			}

			return h(l.WithContext(ctx), method, input, output)
		}
	})
	vars.SvcClient = libsvc.DecorateClient(vars.SvcClient, func(h libsvc.ServiceHandler) libsvc.ServiceHandler {
		return func(ctx context.Context, method libsvc.Method, input, output interface{}) error {
			// Context 中若有 reqID，则传递过去
			reqID := middleware.GetReqID(ctx)
			if reqID != "" {
				ctx = libsvc.WithPassthru(ctx, map[string]string{vars.ReqIDFieldName: reqID})
			}
			return h(ctx, method, input, output)
		}
	})

	return nil
}
