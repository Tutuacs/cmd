package main

import (
	"log"

	"github.com/Tutuacs/pkg/config"
	"github.com/Tutuacs/pkg/db"

	"github.com/Tutuacs/cmd/api"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func main() {

	conf_API := config.GetAPI()

	// TODO: Set the db connection
	dbConnection, err := db.NewConnection()

	if err != nil {
		log.Fatal(string(Red), err, Reset)
	}

	server := api.NewApiServer(conf_API.Port, dbConnection)
	if err := server.Run(); err != nil {
		log.Fatal(string(Red), err, Reset)
	}

}
