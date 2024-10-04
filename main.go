package main

import (
	"fmt"

	"github.com/Tutuacs/pkg/config"
	"github.com/Tutuacs/pkg/db"
	"github.com/Tutuacs/pkg/logs"

	"github.com/Tutuacs/cmd/api"
)

func main() {

	conf_API := config.GetAPI()

	// TODO: Set the db connection
	dbConnection, err := db.NewConnection()

	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error connecting with db: %s", err))
	}

	server := api.NewApiServer(conf_API.Port, dbConnection)
	if err := server.Run(); err != nil {
		logs.ErrorLog(fmt.Sprintf("Error starting server: %s", err))
	}

}
