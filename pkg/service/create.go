// The hierarchy is as follows:
// A Progo build corresponds to a single Docker container
// A build can have multiple steps or tasks (install all packages needed for
// testing, seed the database, run all tests, etc).
// A task can have multiple commands which are basically OS processes (cd, ls,
// apt-get install, etc.).

package service

import (
	"context"

	"github.com/docker/docker/client"

	"progo/build/pkg/docker"
	"progo/build/pkg/entity"
	"progo/build/pkg/log"
)

// Create stores a new build entity.
func (b *buildService) Create(ctx context.Context,
	build *entity.Build) (string, error) {

	cli, err := client.NewClientWithOpts(client.WithVersion("1.40"))
	if err != nil {
		return "", err
	}

	container, err := docker.NewContainer(ctx, cli, build)
	if err != nil {
		log.Print("error", "Error creating new container", err)
		return "", err
	}

	// Runs every instruction in a goroutine.
	go runAll(cli, container, build)

	return container.ID, nil
}

func runAll(cli docker.Client, container *entity.Container,
	build *entity.Build) {

	err := RunTasks(context.Background(), cli, container, []string{})
	if err != nil {
		log.Print("error", "Error running tasks", err)

		return
	}
}
