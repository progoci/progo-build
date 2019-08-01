package service

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"progo/build/pkg/entity"
)

// Create stores a new build entity.
func (c *buildService) Create(ctx context.Context,
	build entity.Build) (string, error) {

	id, _ := createContainer(ctx, build)

	return id, nil
}

func createContainer(ctx context.Context, build entity.Build) (string, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	log.Print(containers)
	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	return "1", nil
}
