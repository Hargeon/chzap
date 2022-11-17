package main

import (
	"context"

	"github.com/Hargeon/chzap/internal/logger"
	"github.com/Hargeon/chzap/internal/repository"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.uber.org/zap"
)

func main() {
	if err := setup(); err != nil {
		panic(err)
	}
}

func setup() error {
	url := "mongodb://mongoadmin:supersecretpasswordlol@localhost:27016"
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(url))
	if err != nil {
		return err
	}

	defer client.Disconnect(context.Background())

	if err := client.Ping(context.Background(), readpref.Primary()); err != nil {
		return err
	}

	col := client.Database("Errors").Collection("Fatal")
	repo := repository.NewRepository(col)

	zLogger := logger.NewLogger(repo)

	zLogger.Info("Some err",
		zap.String("CHeck str", "IAMASTRING"),
		zap.Int64("CHECK int", 56),
	)

	return nil
}
