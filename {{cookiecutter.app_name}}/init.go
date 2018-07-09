package {{ cookiecutter.app_name }}

import (
	"%%baseimport%%/vars"
	"%%baseimport%%/endpoints"
	"github.com/go-chi/chi"
)

// Initialize 初始化该 app
func Initialize() (chi.Router, error) {
	// 初始化全局变量
	// TODO ...

	// 初始化 endpoints
	if err := endpoints.InitEndpoints(); err != nil {
		return nil, err
	}

	return vars.Router, nil
}

