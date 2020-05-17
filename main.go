package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/hzhyvinskyi/micro/handlers"
)

func main() {
	l := log.New(os.Stdout, "api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/goodbye", goodbyeHandler)

	srv := http.Server{
		Addr: ":3000",
		Handler: serveMux,
		IdleTimeout: 120 * time.Second,
	}

	go func () {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatalln(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	log.Printf("Received %s signal. Gracefully shutdown...\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	srv.Shutdown(ctx)
}
