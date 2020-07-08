package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/progoci/progo-core/uuid"
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

// GenerateContainerName creates a new name for a container.
func (Client) GenerateContainerName() string {
	return uuid.Get()
}

// ContainerConfig generates a configuration for creating a container.
func (Client) ContainerConfig(image string, virtualhost string) *container.Config {
	return &container.Config{
		Image: image,
		Env:   []string{"VIRTUAL_HOST=" + virtualhost},
		Tty:   true,
	}
}

// ContainerCreate creates a new container.
func (cli *Client) ContainerCreate(ctx context.Context, opts *CreateOpts) (string, error) {
	res, err := cli.Conn.ContainerCreate(ctx, opts.Config, opts.HostConfig, opts.NetworkingConfig, opts.ContainerName)
	if err != nil {
		return "", err
	}

	return res.ID, nil
}

// ContainerStart starts a container.
func (cli *Client) ContainerStart(ctx context.Context, containerID string,
	options types.ContainerStartOptions) error {

	return cli.Conn.ContainerStart(ctx, containerID, options)
}

// ContainerExecCreate runs a new command in a running container.
func (cli *Client) ContainerExecCreate(ctx context.Context, container string,
	config types.ExecConfig) (types.IDResponse, error) {

	return cli.Conn.ContainerExecCreate(ctx, container, config)
}

// ContainerExecAttach attaches a connection to an exec process in the server
func (cli *Client) ContainerExecAttach(ctx context.Context, execID string,
	config types.ExecConfig) (types.HijackedResponse, error) {

	return cli.Conn.ContainerExecAttach(ctx, execID, config)
}
