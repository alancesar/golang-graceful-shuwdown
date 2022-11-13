package main

import (
	"context"
	"github.com/alancesar/graceful-shuwdown-sample/app"
	"github.com/alancesar/graceful-shuwdown-sample/server"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	engine := gin.Default()
	engine.Handle(http.MethodGet, "/", func(c *gin.Context) {
		log.Println("this request will finish even you press ctrl+c")
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "finished")
	})

	s := server.New(engine, ":8080")
	a := app.New(s)
	a.Start(context.Background())
}
