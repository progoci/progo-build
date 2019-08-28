package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"progo/core/config"
)

// StartConnection creates connection to a MongoDB database.
func StartConnection(db string) Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := "mongodb://" + config.Get("DB_HOST") + ":" + config.Get("DB_PORT")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	return client.Database(db)
}
