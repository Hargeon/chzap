package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *Repository {
	r := &Repository{
		db: db,
	}

	return r
}

func (repo *Repository) Write(p []byte) (int, error) {
	var bdoc interface{}

	if err := bson.UnmarshalExtJSON(p, true, &bdoc); err != nil {
		return 0, err
	}

	if _, err := repo.db.InsertOne(context.Background(), bdoc); err != nil {
		return 0, err
	}

	return len(p), nil
}
