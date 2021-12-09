package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"movie/config"
)

var (
	dbClient *mongo.Database
	client   *mongo.Client
)

const (
	CollectionMovie   = "movie"
	CollectionRating  = "rating"
	CollectionTagUser = "tag_user"
)

func GetClient() *mongo.Database {
	return dbClient
}

func InitModel() error {
	cfg := config.GetConfig().Mongo
	clientOps := options.Client().ApplyURI(cfg.Url)
	clientOps.Auth = &options.Credential{
		Username: cfg.User,
		Password: cfg.Password,
	}
	cli, err := mongo.Connect(context.Background(), clientOps)
	if err != nil {
		return err
	}

	if err = cli.Ping(context.Background(), readpref.Primary()); err != nil {
		return err
	}

	client = cli
	dbClient = cli.Database(cfg.DBName)
	return nil
}

func Disconnect() error {
	return client.Disconnect(context.Background())
}
