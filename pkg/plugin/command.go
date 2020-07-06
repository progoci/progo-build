package plugin

import (
	"context"

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
		runCommand(ctx, opts, command)
	}
}

// runCommand executes a single linux command.
func runCommand(ctx context.Context, opts *RunOpts, cmd string) error {
	// Creates the instance of the process to run.
	exec, err := opts.DockerClient.ContainerExecCreate(ctx, opts.ContainerID, types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          []string{"/bin/bash", "-c", cmd},
		Privileged:   true,
	})
	if err != nil {
		return errors.Wrapf(err, "could not run command %s", cmd)
	}

	// Attach to exec instance's stdout and stderr.
	response, err := opts.DockerClient.ContainerExecAttach(ctx, exec.ID, types.ExecConfig{})
	if err != nil {
		return errors.Wrapf(err, "could not attach to exec instance for command %s", cmd)
	}

	sendOpts := &progolog.SendOpts{
		BuildID:     opts.BuildID,
		ServiceName: opts.ServiceName,
		StepName:    opts.Step.Name,
		StepNumber:  opts.StepNumber,
		Command:     cmd,
	}

	progolog.Send(response.Reader, sendOpts)

	return nil
}
