package main

import (
	"fmt"
	"go-lang-server/config"
	"go-lang-server/router"
	"go-lang-server/util"
	"log"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var cfg *config.GeneralConfig

	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Println(err)
	}

	if dbErr := util.ConnectDB(cfg.DBConfig); dbErr != nil {
		return dbErr
	}

	router.NewRouter().Start()

	return nil
}
