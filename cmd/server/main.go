package main

import (
	"github.com/arzanico/event-processor/internal/server"
	"log"
)

func main() {
	srv := server.NewHttpServer(":8080")
	log.Println("Server is running ...")
	log.Fatal(srv.ListenAndServe())
}
