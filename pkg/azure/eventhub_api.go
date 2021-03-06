package azure

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-event-hubs-go/v3"
	"time"
)

type EventHubApi struct {
	azConn string
}

func(eventHubApi *EventHubApi) Init(azConn string) *EventHubApi {
	eventHubApi.azConn = azConn

	return eventHubApi
}

func(eventHubApi *EventHubApi) Send(data map[string]interface{}) error {

	hub, err := eventhub.NewHubFromConnectionString(eventHubApi.azConn)
	if err != nil { return err }

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	json, err := json.Marshal(data)
	if err != nil { return err }

	err = hub.Send(ctx, eventhub.NewEvent(json))
	if err != nil { return err }

	return nil

}
