package main

import (
	"l0/cache"
	"l0/db"
	"l0/server"
	"l0/sub"
	"log"

	_ "github.com/lib/pq"
)

// @title l0 API
// @version 1.0
// @description API Server

// @host localhost:8000
// @BasePath /
func main() {

	database, err := db.NewPostgresDB()

	if err != nil {
		log.Println(err)
		return
	}

	repo := db.NewRepository(database)

	lcache := cache.NewLocalCache(repo)

	lcache.Initialize()

	handler := server.NewHandler(repo, lcache)

	subscriber, err := sub.NewSubscriber(repo, handler)

	if err != nil {
		log.Println(err)
		return
	}

	subscriber.Subscribe()

	srv := new(server.Server)

	if err = srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal(err.Error())
	}
}
