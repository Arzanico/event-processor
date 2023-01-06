package modules

import (
	"fmt"
	"github.com/arzanico/event-processor/internal/server"
)

const eventType = "realtimeRequest"

func Process(event server.Event) (string, error) {
	fmt.Println(event.Id)
	fmt.Println(eventType)
	return *event.Type, nil
}
