package {{ cookiecutter.app_name }}

import (
	"database/sql"

	"%%baseimport%%/vars"
	_ "github.com/go-sql-driver/mysql"
)

// Initialize 初始化该 app，参数为:
//  - urlPrefix: 该 app endpoints 的 url 前缀
//  - dsn: 数据库 dsn
func Initialize(
	urlPrefix string,
	dsn string,
) error {
	// url 前缀
	vars.URLPrefix = urlPrefix

	// 数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	vars.DSN = dsn
	vars.DB = db

	return nil
}
