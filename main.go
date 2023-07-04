package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	mux := chi.NewMux()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("this request will finish even you press ctrl+c")
		time.Sleep(5 * time.Second)
		_, _ = w.Write([]byte("finished"))
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	log.Println("all systems go!")

	<-ctx.Done()
	stop()

	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println(err)
	}

	log.Println("good bye")
}
