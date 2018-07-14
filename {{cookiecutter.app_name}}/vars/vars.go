package vars

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/rs/zerolog"
)

// 全局变量，必须首先初始化它们
var (
	// Router 是本 app endpoints 的路由器
	Router chi.Router = chi.NewRouter()

	// 日志
	Logger zerolog.Logger

	// 数据库
	DB *sql.DB

	// TODO ...
)
