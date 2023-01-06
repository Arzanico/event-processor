package eventprocessor

import (
	"fmt"
	"github.com/arzanico/event-processor/internal/server"
)

type EventProcessor interface {
	ProcessEvent(event server.Event) (string, error)
}

type Processor struct {
	EventProcessor EventProcessor
}

func (p Processor) Process(event server.Event) (string, error) {
	return p.EventProcessor.ProcessEvent(event)
}

type RealTimeRequest struct{}

func (r RealTimeRequest) Process(event server.Event) (string, error) {
	fmt.Println("Define RealTime behavior")
	return "done", nil
}

func Process(event server.Event) (string, error) {
	fmt.Println(event.Id)
	return *event.Type, nil
}
