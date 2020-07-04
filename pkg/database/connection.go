package database

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Opts is the information to connnect to the database.
// If URI is set, it overwrites all other connection configuration.
type Opts struct {
	Host     string
	Port     string
	Database string
	URI      string
}

// StartConnection creates connection to a MongoDB database.
func StartConnection(opts *Opts) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	uri := opts.URI
	if uri == "" {
		uri = "mongodb://" + opts.Host + ":" + opts.Port
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	return client.Database(opts.Database), nil
}
