package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/terachanple/redis-ranking-demo/config"
	"github.com/terachanple/redis-ranking-demo/handler"
	"github.com/terachanple/redis-ranking-demo/service"
)

func main() {
	config := config.New("")

	service.Initialize(config)

	router := httprouter.New()
	router.GET("/rankings", handler.Index)
	router.GET("/rankings/:id", handler.Show)
	router.POST("/rankings", handler.CreateWeeklyRanking)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), router))
}
