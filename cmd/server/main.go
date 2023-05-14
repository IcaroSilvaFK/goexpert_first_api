package main

import (
	"log"
	"net/http"

	"github.com/IcaroSilvaFK/goexpert_first_api/configs"
)

func main() {

	cfg, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(cfg.WebServerPort, nil)
}
