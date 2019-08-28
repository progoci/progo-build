// The hierarchy is as follows:
// A Progo build corresponds to a single Docker container
// A build can have multiple steps or tasks (install all packages needed for
// testing, seed the database, run all tests, etc).
// A task can have multiple commands which are basically OS processes (cd, ls,
// apt-get install, etc.).

package docker

import (
	"context"
	"errors"

	dockerTypes "github.com/docker/docker/api/types"

	"progo/build/pkg/entity"
	"progo/build/pkg/types"
	"progo/core/log"
)

// NewContainer creates a new Docker container and runs it.
func NewContainer(ctx context.Context, cli Client,
	build *entity.Build) (*entity.Container, error) {

	if _, ok := availableImages[build.Image]; !ok {
		return nil, errors.New("image is not valid")
	}

	newContainer, err := createContainer(ctx, cli, build)
	if err != nil {
		return nil, err
	}

	// Runs the container.
	err = cli.ContainerStart(ctx, newContainer.ID, dockerTypes.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}

	log.Print("info", "", "Container running at "+newContainer.Host)

	return newContainer, nil
}

// RunTasks runs all the steps set in the config file.
func RunTasks(ctx context.Context, cli Client, loomConn types.LoomSocket,
	container *entity.Container, build *entity.Build) error {

	// Creates a new BuildLogs.
	logs := &types.BuildLogs{BuildID: build.ID}

	for _, task := range build.Tasks {
		return runTask(ctx, cli, logs, container, &task)
	}

	return nil
}
