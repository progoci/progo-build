// The hierarchy is as follows:
// A Progo build corresponds to a single Docker container
// A build can have multiple steps or tasks (install all packages needed for
// testing, seed the database, run all tests, etc).
// A task can have multiple commands which are basically OS processes (cd, ls,
// apt-get install, etc.).

package docker

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	dockerClient "github.com/docker/docker/client"
	log "github.com/sirupsen/logrus"

	"github.com/progoci/progo-build/internal/entity"
)

// Docker describes a handler for docker-related tasks.
type Docker interface {
	ContainerCreate(ctx context.Context, opts *CreateOpts) (string, error)
	ContainerConfig(image string, virtualhost string) *container.Config
	ContainerStart(ctx context.Context, containerID string, options types.ContainerStartOptions) error

	NetworkConnect(ctx context.Context, preffix string, name string) (*network.NetworkingConfig, error)
	NetworkingConfig(networkID string) *network.NetworkingConfig
}

// Client handles all docker-related tasks (implementation of Docker interface).
type Client struct {
	Conn             *dockerClient.Client
	Logger           *log.Logger
	ProxyContainerID string
}

// New creates a new docker client.
func New(log *log.Logger, proxy string) (*Client, error) {
	client, err := dockerClient.NewEnvClient()
	if err != nil {
		return nil, err
	}

	return &Client{
		Conn:             client,
		Logger:           log,
		ProxyContainerID: proxy,
	}, nil
}

// RunTasks runs all the steps set in the config file.
func RunTasks(ctx context.Context, cli Client, container *Container, build *entity.Build) error {

	// Creates a new BuildLogs.
	//logs := &types.BuildLogs{BuildID: build.ID}

	for _, task := range build.Tasks {
		return runTask(ctx, cli, container, &task)
	}

	return nil
}
