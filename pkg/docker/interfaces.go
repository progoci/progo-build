package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

// Client describes a Docker client.
type Client interface {
	ContainerCreate(ctx context.Context, config *container.Config,
		hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig,
		containerName string) (container.ContainerCreateCreatedBody, error)

	ContainerExecCreate(ctx context.Context, container string,
		config types.ExecConfig) (types.IDResponse, error)

	ContainerExecStart(ctx context.Context, execID string,
		config types.ExecStartCheck) error

	ContainerStart(ctx context.Context, containerID string,
		options types.ContainerStartOptions) error
}
