package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	dockerClient "github.com/docker/docker/client"
)

// Docker describes a handler for docker-related tasks.
type Docker interface {
	GenerateContainerName() string
	ContainerCreate(ctx context.Context, opts *CreateOpts) (string, error)
	ContainerConfig(image string, virtualhost string) *container.Config
	ContainerStart(ctx context.Context, containerID string, options dockerTypes.ContainerStartOptions) error
	ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error)
	ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error)

	NetworkConnect(ctx context.Context, preffix string, name string) (*network.NetworkingConfig, error)
	NetworkingConfig(networkID string) *network.NetworkingConfig
}

// Client handles all docker-related tasks (implementation of Docker interface).
type Client struct {
	Conn             *dockerClient.Client
	ProxyContainerID string
}

// New creates a new docker client.
func New(proxier string) (*Client, error) {
	client, err := dockerClient.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		Conn:             client,
		ProxyContainerID: proxier,
	}, nil
}
