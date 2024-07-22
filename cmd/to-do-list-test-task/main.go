package main

import (
	log "github.com/sirupsen/logrus"
	"to-do-list-test-task/config"
	"to-do-list-test-task/service/http"
	"to-do-list-test-task/storage/postgre"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	pClient, err := postgre.New(&cfg.PostgresConfig)
	if err != nil {
		log.Fatal(err)
	}

	svr := http.New(&cfg.HttpConfig, pClient)
	err = svr.Run()
	if err != nil {
		log.Fatal(err)
	}
}
