package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(addr string) *http.Server {
	httpsrv := newHttpServer()
	r := mux.NewRouter()
	r.HandleFunc("/realtimerequest", httpsrv.realtimerequest).Methods(http.MethodPost)
	r.HandleFunc("/pushnotification", httpsrv.pushnotification).Methods(http.MethodPost)
	r.HandleFunc("/initialload", httpsrv.initialload).Methods(http.MethodPost)
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

type httpServer struct {
	Event *Event
}

func newHttpServer() *httpServer {
	return &httpServer{
		Event: NewEvent(),
	}
}

func (s *httpServer) realtimerequest(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	thisEvent := Event{
		Id:   event.Id,
		Type: event.Type,
	}

	newEvent := Processor{RealTimeRequest{}}
	newEvent.Process(thisEvent)

	res := thisEvent
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *httpServer) pushnotification(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	thisEvent := Event{
		Id:   event.Id,
		Type: event.Type,
	}

	newEvent := Processor{PushNotification{}}
	newEvent.Process(thisEvent)

	res := thisEvent
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (s *httpServer) initialload(w http.ResponseWriter, r *http.Request) {
	var event Event
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	thisEvent := Event{
		Id:   event.Id,
		Type: event.Type,
	}

	newEvent := Processor{InitialUpload{}}
	newEvent.Process(thisEvent)

	res := thisEvent
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
