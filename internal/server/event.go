package server

type Event struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

func NewEvent() *Event {
	return &Event{}
}
