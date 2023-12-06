package handlers

import (
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/config"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/db"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/logger"
)

type Backend struct {
	log *logger.Logger
	db  *db.Mongo
}

func NewBackend(cfg *config.Config, mongodb *db.Mongo, log *logger.Logger) (*Backend, error) {
	return &Backend{
		log: log,
		db:  mongodb,
	}, nil
}
