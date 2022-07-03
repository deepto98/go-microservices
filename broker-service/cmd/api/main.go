package main

import (
	"fmt"
	"log"
	"net/http"
)

//Broker which helps services to
// do inter-communication via messaging is Message Broker.

//Note to self -
// got bind: permission denied error when trying 80
// got missing port in address when just specifying port no

const webPort = "80"

type Config struct{}

func main() {
	app := Config{}

	log.Println("Starting broker service on port %s", webPort)

	//define http server, using handlers from routes.go
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Panic((err))
	}
}
