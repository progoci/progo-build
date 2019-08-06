package service

import (
	"context"
	"log"

	"github.com/docker/docker/client"

	"progo/build/pkg/docker"
	"progo/build/pkg/entity"
)

// Create stores a new build entity.
func (c *buildService) Create(ctx context.Context,
	build entity.Build) (string, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}

	id, err := docker.NewContainer(ctx, cli, &build)
	if err != nil {
		log.Print(err)
		return "", err
	}

	return id, nil
}
