package plugin

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"

	"github.com/progoci/progo-build/pkg/docker"
)

// Command is a plugin for runinng commands in the shell.
type Command struct {
	docker docker.Docker
}

// NewCommand initializes an instance of the plugin.
func NewCommand() *Command {
	return &Command{}
}

// Run executes a command instructions for a given step.
func (Command) Run(ctx context.Context, opts *RunOpts) {
	for _, command := range opts.Step.Commands {
		cmd := []string{"/bin/bash", "-c", command}

		runCommand(ctx, opts, cmd)
	}
}

// runCommand executes a single linux command.
func runCommand(ctx context.Context, opts *RunOpts, cmd []string) error {
	// Creates the instance of the process to run.
	exec, err := opts.DockerClient.ContainerExecCreate(ctx, opts.ContainerID, types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
		Privileged:   true,
	})
	if err != nil {
		msg := fmt.Sprintf("could not run command %s", strings.Join(cmd, " "))
		return errors.Wrap(err, msg)
	}

	// Attach to exec instance's stdout and stderr.
	response, err := opts.DockerClient.ContainerExecAttach(ctx, exec.ID, types.ExecConfig{})
	if err != nil {
		msg := fmt.Sprintf("could not attach to exec instance for command %s", strings.Join(cmd, " "))
		return errors.Wrap(err, msg)
	}

	io.Copy(os.Stdout, response.Reader)

	return nil
}
