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

func (r RealTimeRequest) ProcessEvent(event server.Event) (string, error) {
	fmt.Println("Define RealTime behavior")
	return "done", nil
}

type PushNotification struct{}

func (p PushNotification) ProcessEvent(event server.Event) (string, error) {
	fmt.Println("Define PushNotification behavior")
	return "done", nil
}

type InitialUpload struct{}

func (i InitialUpload) ProcessEvent(event server.Event) (string, error) {
	fmt.Println("Define InitialLoad behavior")
	return "done", nil
}
