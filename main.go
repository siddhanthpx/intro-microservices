package main

import (
	"intro-microservices/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "intro", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":9090", sm)
}
