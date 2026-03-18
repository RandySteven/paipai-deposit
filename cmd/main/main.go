package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RandySteven/go-kopi/apps"
	"github.com/RandySteven/go-kopi/configs"
	"github.com/RandySteven/go-kopi/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("./files/env/.env")
	if err != nil {
		log.Fatalln(`failed to load .env `, err)
		return
	}
}

func main() {
	configPath, err := configs.ParseFlags()
	if err != nil {
		log.Fatalln(err)
		return
	}

	config, err := configs.NewConfig(configPath)
	if err != nil {
		log.Fatalln(err)
		return
	}
	ctx := context.TODO()

	app, err := apps.NewApp(config)
	if err != nil {
		log.Fatalln(`Error starting app `, err)
		return
	}

	apis := app.PrepareHttpHandler(ctx)
	r := mux.NewRouter()
	router := routes.NewEndpointRouters(apis)
	routes.InitRouter(router, r)

	if err = app.Temporal.Start(); err != nil {
		log.Fatalln("Failed to start Temporal worker:", err)
		return
	}
	defer app.Temporal.Stop()

	go config.Run(r)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = app.RefreshRedis(ctx); err != nil {
		log.Fatal(err)
		return
	}
}
