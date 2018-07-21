package {{ cookiecutter.app_name }}

import (
	"database/sql"

	"%%baseimport%%/endpoints"
	"%%baseimport%%/services"
	"%%baseimport%%/vars"
	"bitbucket.org/jayven/platform-kit/svc"
	"github.com/go-chi/chi"
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
	vars.SvcServer = svcServer
	vars.SvcClient = svcClient

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

