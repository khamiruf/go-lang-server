package main

import (
	"fmt"
	"go-lang-server/router"
	"go-lang-server/util"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	if dbErr := util.ConnectDB(); dbErr != nil {
		return dbErr
	}

	router.NewRouter()

	return nil
}
