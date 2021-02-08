package azure

import (
	"context"
	"errors"
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

	if data["ts"] != nil { return errors.New("missing [ts] field")}

	if data["device_id"] != nil { return errors.New("missing [device_id] field")}


	hub, err := eventhub.NewHubFromConnectionString(eventHubApi.azConn)
	if err != nil { return err }

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err = hub.Send(ctx, eventhub.NewEvent(data))
	if err != nil { return err }

	return nil

}
