package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBRepository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewMongoDBUserRepository(uri, database, collection string) (*MongoDBRepository, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	db := client.Database(database)
	coll := db.Collection(collection)

	return &MongoDBRepository{
		client:     client,
		collection: coll,
	}, nil
}
