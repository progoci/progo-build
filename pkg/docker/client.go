package docker

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"

	"progo/build/pkg/entity"
	"progo/build/pkg/log"
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
	err = cli.ContainerStart(ctx, newContainer.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, err
	}

	log.Print("info", "", "Container running at "+newContainer.Host)

	return newContainer, nil
}

// RunTasks run the setup tasks in config file.
func RunTasks(ctx context.Context, cli Client, container *entity.Container,
	cmds []string) error {

	exec, err := cli.ContainerExecCreate(ctx, container.ID, types.ExecConfig{
		Cmd:        cmds,
		Privileged: true,
		Detach:     true,
	})
	if err != nil {
		log.Print("error", "Error creating exec instance", err)
		return err
	}

	err = cli.ContainerExecStart(ctx, exec.ID, types.ExecStartCheck{
		Detach: true,
	})
	if err != nil {
		log.Print("error", "Error starting exec instance", err)
		return err
	}

	log.Print("info", "Exec info", exec)

	return nil
}
