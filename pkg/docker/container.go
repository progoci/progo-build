package docker

import (
	"context"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"

	"progo/build/config"
	"progo/build/pkg/entity"
	"progo/build/pkg/utils"
)

func createContainer(ctx context.Context, cli Client,
	build *entity.Build) (*entity.Container, error) {

	uuid := utils.GetUUID()
	host := uuid + "." + config.Get("DEFAULT_HOST")
	networkID := config.Get("DOCKER_NETWORK")

	containerConfig := &container.Config{
		Image: build.Image,
		Env:   []string{"VIRTUAL_HOST=" + host},
		Tty:   true,
	}

	networkConfig := &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"net": &network.EndpointSettings{
				NetworkID: networkID,
			},
		},
	}

	// Creates a new container.
	newContainer, err := cli.ContainerCreate(ctx, containerConfig, nil, networkConfig, "")
	if err != nil {
		return nil, err
	}

	return &entity.Container{
		ID:   newContainer.ID,
		Host: host,
	}, nil
}
