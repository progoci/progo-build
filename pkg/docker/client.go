package docker

import (
	"context"
	"errors"

	"github.com/docker/docker/api/types"

	"progo/build/pkg/entity"
)

// NewContainer creates a new Docker container and runs it.
func NewContainer(ctx context.Context, cli Client, build *entity.Build) (string, error) {

	if _, ok := availableImages[build.Image]; !ok {
		return "", errors.New("Image is not valid")
	}

	containerID, err := createContainer(ctx, cli, build)
	if err != nil {
		return "", err
	}

	// Runs the container.
	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return "", err
	}

	return "1", nil
}
