package {{ cookiecutter.app_name }}

import (
	"database/sql"

	"%%baseimport%%/endpoints"
	"%%baseimport%%/services"
	"%%baseimport%%/vars"
	"bitbucket.org/jayven/platform-kit/svc"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/zerolog"
)

// Initialize 初始化该 app
func Initialize(
	logger *zerolog.Logger,
	db *sql.DB,
	svcServer libsvc.ServiceServer,
	svcClient libsvc.ServiceClient,
) (chi.Router, error) {
	// 日志
	vars.Logger = logger.With().Timestamp().Str("app", "{{cookiecutter.app_name}}").Logger()

	// 数据库
	vars.DB = db

	// 服务
	vars.SvcServer = libsvc.DecorateServer(svcServer, func(h libsvc.ServiceHandler) libsvc.ServiceHandler {
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
	vars.SvcClient = libsvc.DecorateClient(svcClient, func(h libsvc.ServiceHandler) libsvc.ServiceHandler {
		return func(ctx context.Context, method libsvc.Method, input, output interface{}) error {
			// Context 中若有 reqID，则传递过去
			reqID := middleware.GetReqID(ctx)
			if reqID != "" {
				ctx = libsvc.WithPassthru(ctx, map[string]string{vars.ReqIDFieldName: reqID})
			}
			return h(ctx, method, input, output)
		}
	})

	// TODO 初始化其它全局变量
	// ...

	// 初始化 services
	if err := services.InitServices(); err != nil {
		return nil, err
	}

	// 初始化 endpoints
	if err := endpoints.InitEndpoints(); err != nil {
		return nil, err
	}

	return vars.Router, nil
}

