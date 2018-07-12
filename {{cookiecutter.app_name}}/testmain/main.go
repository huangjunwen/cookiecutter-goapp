package main

import (
	"log"
	"net/http"

	"%%baseimport%%"
	"github.com/rs/zerolog"
)

func main() {
	rt, err := {{cookiecutter.app_name}}.Initialize(
		"root:123456@tcp(localhost:{{cookiecutter.devdb_port}})/dev?parseTime=true",
		zerolog.InfoLevel,
	)
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":3333", rt)
	
}
