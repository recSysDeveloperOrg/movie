package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"movie/config"
	"time"
)

var dbClient *mongo.Database

const (
	CollectionMovie = "movie"
	CollectionRating = "rating"
	CollectionTagUser = "tag_user"
)

func GetClient() *mongo.Database {
	return dbClient
}

func InitModel() error {
	cfg := config.GetConfig()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.MaxConnectionTimeout)*time.Second)
	defer cancel()
	clientOps := options.Client().ApplyURI(cfg.Url)
	clientOps.Auth.Username, clientOps.Auth.Password = cfg.User, cfg.Password
	cli, err := mongo.Connect(ctx, clientOps)
	if err != nil {
		return err
	}

	dbClient = cli.Database(cfg.DBName)
	return nil
}
