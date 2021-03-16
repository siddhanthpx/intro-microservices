package main

import (
	"context"
	"intro-microservices/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const TIMEOUT = 30 * time.Second

func main() {
	l := log.New(os.Stdout, "intro", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	toc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(toc)
}
