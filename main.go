package main

import (
	"fmt"

	"github.com/Tutuacs/pkg/config"
	"github.com/Tutuacs/pkg/logs"

	"github.com/Tutuacs/cmd/api"
)

func main() {
	conf_API := config.GetAPI()

	server, err := api.NewApiServer(conf_API.Port)
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error creating server: %s", err))
		return
	}

	// ! Want to use MQTT? Uncomment the following lines
	// ! You can init on main.go or on api.go
	// * Recomended not init on both files
	// * Create hanldersFunctions on the pkg/mqtt package
	// ? You can "UseMQTT() inside http:handlers too"
	// mqttClient, err := mqtt.UseMqtt()
	// if err != nil {
	// 	logs.ErrorLog(fmt.Sprintf("Error connecting Mqtt: %s", err))
	// }

	// ! Want to use Redis pub/sub? Uncomment the following lines
	// ! You can init on main.go or on api.go
	// * Recomended not init on both files
	// * Create hanldersFunctions on the pkg/cache package
	// pubSub, err := pubsub.UsePubSubService()
	// if err != nil {
	// 	logs.ErrorLog(fmt.Sprintf("Error creating PubSubService: %s", err))
	// }
	// pubSub.Run()

	if err := server.Run(); err != nil {
		logs.ErrorLog(fmt.Sprintf("Error starting server: %s", err))
	}
}
