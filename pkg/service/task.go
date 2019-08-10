package service

import (
	"context"

	"progo/build/pkg/docker"
	"progo/build/pkg/entity"
)

// RunTasks starts the steps set in the config file.
func RunTasks(ctx context.Context, cli docker.Client,
	container *entity.Container, tasks []string) error {

	tasks = []string{"/bin/bash", "-c", "cat /var/log/bootstrap.log"}

	return docker.RunTasks(ctx, cli, container, tasks)
}
