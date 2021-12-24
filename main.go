package main

import (
	"log"

	"github.com/EgorSkurihin/Hokku/api"
	"github.com/EgorSkurihin/Hokku/config"
	_ "github.com/EgorSkurihin/Hokku/docs"
	"github.com/EgorSkurihin/Hokku/store/mysql_store"
)

// @title Hokku Rest API
// @This is a sample server.
// @version 1.0

// @host localhost:1323
// @BasePath /
// @schemes http

// @securityDefinitions.apiKey cookieAuth
// @in cookie
// @name session
func main() {
	//Read config file
	conf, err := config.New("config/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	// Create and open storage
	store := mysql_store.New(&conf.Store)
	if err := store.Open(); err != nil {
		log.Fatal(err)
	}
	defer store.Close()

	// Start API Server
	api := api.New(&conf.Server, store)
	log.Fatal(api.Start())
}
