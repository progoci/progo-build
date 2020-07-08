package database

import (
	"context"
	"time"

	"github.com/pkg/errors"
	"github.com/progoci/progo-build/pkg/build"
)

// Create inserts a new build.
func (d *Database) Create(build *build.Build) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := d.MongoClient.Collection("builds").InsertOne(ctx, build)
	if err != nil {
		return "", errors.Wrap(err, "could not insert build")
	}

	return result.InsertedID.(string), nil
}
