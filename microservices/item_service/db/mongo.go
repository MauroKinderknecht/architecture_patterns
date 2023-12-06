package db

import (
	"context"
	"time"

	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/config"
	"github.com/MauroKinderknecht/tech_talk/microservices/item_service/logger"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	MongoDatabaseName = "backend"
)

var (
	ErrorNotFound = errors.New("Not Found")
)

type Mongo struct {
	Client      *mongo.Client
	DB          *mongo.Database
	collections map[string]*mongo.Collection
	log         *logger.Logger
}

func NewMongo(config *config.Config, log *logger.Logger) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	bsonOpts := &options.BSONOptions{
		UseJSONStructTags:       true,
		NilMapAsEmpty:           true,
		NilSliceAsEmpty:         true,
		StringifyMapKeysWithFmt: true,
	}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().
		SetBSONOptions(bsonOpts).
		ApplyURI(config.MongoURL).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	log = log.With(logger.String("component", "mongodb"))

	log.Info("created client, about to ping db")
	db := client.Database(MongoDatabaseName)
	if err := db.RunCommand(ctx, bson.D{{Key: "ping", Value: 1}}).Err(); err != nil {
		return nil, errors.Wrap(err, "failed to ping db")
	}
	log.Info("pinged db successfully")

	return &Mongo{DB: db, Client: client, log: log, collections: make(map[string]*mongo.Collection)}, nil
}
