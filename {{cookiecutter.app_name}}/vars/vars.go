package vars

import (
	"github.com/go-chi/chi"
)

// 全局变量，必须首先初始化它们
var (
	// Router 是本 app endpoints 的路由器
	Router chi.Router = chi.NewRouter()

	// TODO ...
)
