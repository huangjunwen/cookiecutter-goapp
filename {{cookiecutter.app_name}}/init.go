package {{ cookiecutter.app_name }}

import (
	"database/sql"

	"%%baseimport%%/endpoints"
	"%%baseimport%%/vars"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

// Initialize 初始化该 app
func Initialize(
	db *sql.DB,
	logger *zerolog.Logger,
) (chi.Router, error) {
	// 日志
	vars.Logger = logger.With().Timestamp().Str("app", "{{cookiecutter.app_name}}").Logger()

	// 数据库
	vars.DB = db

	// TODO 初始化其它全局变量
	// ...

	// 初始化 endpoints
	if err := endpoints.InitEndpoints(); err != nil {
		return nil, err
	}

	return vars.Router, nil
}

