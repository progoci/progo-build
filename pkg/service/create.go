package service

import (
	"context"

	"github.com/docker/docker/client"

	"progo/build/pkg/docker"
	"progo/build/pkg/entity"
	"progo/build/pkg/log"
)

// Create stores a new build entity.
func (c *buildService) Create(ctx context.Context,
	build entity.Build) (string, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}

	container, err := docker.NewContainer(ctx, cli, &build)
	if err != nil {
		log.Print("error", "Error creating new container", err)
		return "", err
	}

	err = RunTasks(ctx, cli, container, []string{})
	if err != nil {
		log.Print("error", "Error running tasks", err)
		return "", err
	}

	return container.ID, nil
}
