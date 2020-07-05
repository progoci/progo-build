package plugin

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"

	"github.com/progoci/progo-build/pkg/docker"
	"github.com/progoci/progo-build/progolog"
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

	sendOpts := &progolog.SendOpts{
		BuildID:     opts.BuildID,
		ServiceName: opts.ServiceName,
		StepName:    opts.Step.Name,
		StepNumber:  opts.StepNumber,
	}

	progolog.Send(response.Reader, sendOpts)

	return nil
}
