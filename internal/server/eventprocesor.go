package server

import (
	"fmt"
)

type EventProcessor interface {
	ProcessEvent(event Event) (string, error)
}

type Processor struct {
	EventProcessor EventProcessor
}

func (p Processor) Process(event Event) (string, error) {
	return p.EventProcessor.ProcessEvent(event)
}

type RealTimeRequest struct{}

func (r RealTimeRequest) ProcessEvent(event Event) (string, error) {
	fmt.Println("Define RealTime behavior")
	fmt.Printf("Event id : %s\n", event.Id)
	fmt.Printf("Event type : %s\n", event.Type)
	return "done", nil
}

type PushNotification struct{}

func (p PushNotification) ProcessEvent(event Event) (string, error) {
	fmt.Println("Define Pushnotification behavior")
	fmt.Printf("Event id : %s\n", event.Id)
	fmt.Printf("Event type : %s\n", event.Type)
	return "done", nil
}

type InitialUpload struct{}

func (i InitialUpload) ProcessEvent(event Event) (string, error) {
	fmt.Println("Define Initial upload behavior")
	fmt.Printf("Event id : %s\n", event.Id)
	fmt.Printf("Event type : %s\n", event.Type)
	return "done", nil
}
