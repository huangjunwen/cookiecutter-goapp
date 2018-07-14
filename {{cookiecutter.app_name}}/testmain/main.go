package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"%%baseimport%%"
	"github.com/rs/zerolog"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:{{cookiecutter.devdb_port}})/dev?parseTime=true")
	if err != nil {
					return nil, err
	}
	defer db.Close()

	logger := zerolog.New(os.Stdout)

	rt, err := {{cookiecutter.app_name}}.Initialize(
		db,
		&logger,
	)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":3333", rt)
	
}
