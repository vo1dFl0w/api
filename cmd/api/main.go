package main

import (
	"log"

	"github.com/vo1dFl0w/test_api/internal/app/apiserver"
	"github.com/vo1dFl0w/test_api/internal/app/apiserver/config"
)

func main() {
	// TODO run server
	// TODO run cfg
	cfg := config.InitConfig()
	if err := apiserver.Run(cfg); err != nil {
		log.Fatalf("cannot run the server: %v\n", err)
	}

	// TODO run logger
}