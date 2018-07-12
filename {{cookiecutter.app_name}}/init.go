package {{ cookiecutter.app_name }}

import (
	"database/sql"
	"os"

	"%%baseimport%%/endpoints"
	"%%baseimport%%/vars"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

// Initialize 初始化该 app
func Initialize(
	dsn string,
	logLevel zerolog.Level,
) (chi.Router, error) {
	// 日志
	{
		vars.LogLevel = logLevel
		vars.Logger = zerolog.New(os.Stderr).Level(vars.LogLevel).With().Timestamp().Str("app", "{{cookiecutter.app_name}}").Logger()
	}

	// 数据库
	{
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil, err
		}
		vars.DSN = dsn
		vars.DB = db
	}

	// TODO 初始化其它全局变量
	// ...

	// 初始化 endpoints
	if err := endpoints.InitEndpoints(); err != nil {
		return nil, err
	}

	return vars.Router, nil
}

