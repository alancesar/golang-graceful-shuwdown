package main

import (
	"context"
	"github.com/alancesar/graceful-shuwdown-sample/app"
	"github.com/alancesar/graceful-shuwdown-sample/server"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := chi.NewMux()
	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("this request will finish even you press ctrl+c")
		time.Sleep(5 * time.Second)
		_, _ = w.Write([]byte("finished"))
	})

	s := server.New(mux, ":8080")
	a := app.New(s)
	a.Start(context.Background())
}
