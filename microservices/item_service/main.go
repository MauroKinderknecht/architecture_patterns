package main

import (
	"os"
	"os/signal"

	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/config"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/db"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/handlers"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/logger"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/server"
	"github.com/pkg/errors"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(errors.Wrap(err, "config error"))
	}

	log, err := logger.NewLogger(cfg)
	if err != nil {
		panic(errors.Wrap(err, "logger error"))
	}

	mongo, err := db.NewMongo(cfg, log)
	if err != nil {
		panic(errors.Wrap(err, "mongo error"))
	}

	handlers, err := handlers.NewBackend(cfg, mongo, log)
	if err != nil {
		panic(errors.Wrap(err, "backend error"))
	}

	srv, err := server.NewServer(handlers, mongo, log, cfg.ApiPort)
	if err != nil {
		panic(errors.Wrap(err, "server error"))
	}

	if err := srv.Start(); err != nil {
		panic(errors.Wrap(err, "server error"))
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	<-sigs // wait for termination signal

	_ = srv.Shutdown()
}
