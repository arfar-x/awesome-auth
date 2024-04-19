package main

import (
	"fmt"

	"awesome-auth/cmd/http"
	"awesome-auth/configs"
	"awesome-auth/internal/database"
)

var config *configs.AppConfig

func init() {
	config = configs.InitConfig()
	if config == nil {
		panic("Could not initialize configs.")
	}
}

func main() {
	fmt.Println("* Awesome-auth fired up the engines")

	db, err := database.InitDbClient(config)
	if err != nil {
		panic("Could not open connection to the database")
	}

	server := http.Server{Config: config, DB: *db}
	server.Run()
}
