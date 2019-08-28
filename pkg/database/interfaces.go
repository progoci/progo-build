package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client describes a MongoDB client.
type Client interface {
}

// Database describes a connection to a MongoDB database.
type Database interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

// Collection describes a connection to a MongoDB collection.
type Collection interface {
	FindOne(ctx context.Context, filter interface{},
		opts ...*options.FindOneOptions) *mongo.SingleResult

	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)

	UpdateOne(ctx context.Context, filter interface{}, update interface{},
		opts ...*options.UpdateOptions) (*mongo.UpdateResult, error)
}
