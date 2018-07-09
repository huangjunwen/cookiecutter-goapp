package vars

import (
	"database/sql"

	"github.com/go-chi/chi"
)

// 全局变量，必须首先初始化它们
var (
	// Router 是本 app endpoints 的路由器
	Router chi.Router = chi.NewRouter()

	// 数据库连接
	DSN string
	DB  *sql.DB

	// URLPrefix 为本 app endpoints 的 URL 前缀，例如若某 endpoint 为 "/foo/"
	// 而 URLPrefix 为 "https://example.com/prefix/", 则该 endpoint 的完整路径为 "https://example.com/prefix/foo/"
	URLPrefix string
)
