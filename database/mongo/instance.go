package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type I struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func (db *I) Connect(url string, dbName string) (*I, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return nil, err
	}
	db.Client = client

	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, err
	}

	db.DB = client.Database(dbName)
	return db, nil
}
