package {{ cookiecutter.app_name }}

import (
	"os"

	"%%baseimport%%/vars"
	"%%baseimport%%/endpoints"
	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

// Initialize 初始化该 app
func Initialize(
	logLevel zerolog.Level,
) (chi.Router, error) {
	// 初始化日志
	vars.LogLevel = logLevel
	vars.Logger = zerolog.New(os.Stdout).Level(vars.LogLevel).With().Timestamp().Str("app", "{{cookiecutter.app_name}}").Logger()

	// TODO 初始化其它全局变量
	// ...

	// 初始化 endpoints
	if err := endpoints.InitEndpoints(); err != nil {
		return nil, err
	}

	return vars.Router, nil
}

