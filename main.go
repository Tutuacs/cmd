package main

import (
	"fmt"

	"github.com/Tutuacs/pkg/config"
	"github.com/Tutuacs/pkg/logs"
	"github.com/Tutuacs/pkg/mqtt"

	"github.com/Tutuacs/cmd/api"
	pubsub "github.com/Tutuacs/cmd/pub-sub"
)

func main() {
	conf_API := config.GetAPI()

	server, err := api.NewApiServer(conf_API.Port)
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error creating server: %s", err))
		return
	}

	pubSub, err := pubsub.UsePubSubService()
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error creating PubSubService: %s", err))
	}

	mqttClient, err := mqtt.UseMqtt()
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error creating Mqtt: %s", err))
	}

	mqttClient.Subscribe("/hello", mqtt.HandleHello)
	mqttClient.Subscribe("/unknown", mqtt.HandleUnknown)

	pubSub.Run()

	if err := server.Run(); err != nil {
		logs.ErrorLog(fmt.Sprintf("Error starting server: %s", err))
	}

}
