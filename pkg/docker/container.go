package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
)

// Container describes a new container.
type Container struct {
	ID      string // The container ID (assigned at creation by Docker).
	Name    string
	Host    string // The virtualhost where the live site is accessible from (used by nginx-proxy).
	Network string
}

// CreateOpts are the parameter to ContainerCreate
type CreateOpts struct {
	Config           *container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
	ContainerName    string
}

// ContainerConfig generates a configuration for creating a container.
func (client *Client) ContainerConfig(image string, virtualhost string) *container.Config {
	return &container.Config{
		Image: image,
		Env:   []string{"VIRTUAL_HOST=" + virtualhost},
		Tty:   true,
	}
}

// ContainerCreate creates a new container.
func (client *Client) ContainerCreate(ctx context.Context, opts *CreateOpts) (string, error) {
	res, err := client.Conn.ContainerCreate(ctx, opts.Config, opts.HostConfig, opts.NetworkingConfig, opts.ContainerName)
	if err != nil {
		return "", err
	}

	return res.ID, nil
}

// ContainerStart starts a container.
func (client *Client) ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error {
	return client.Conn.ContainerStart(ctx, containerID, options)
}
