package docker

import (
	"context"

	"github.com/docker/docker/api/types"

	"github.com/progoci/progo-build/internal/entity"
)

// buffer is the size of the build logs.
const bufferSize = 1024

// runTask runs a single step in a build.
func runTask(ctx context.Context, cli Client, container *Container, task *entity.Task) error {

	// Each task has its own UUID which is used to retrieve its logs.
	//taskUUID := uuid.Get()
	//logs.TaskUUID = taskUUID

	for _, c := range task.Commands {
		cmd := []string{"/bin/bash", "-c", c}

		runCommand(ctx, cli, container, cmd)
	}

	return nil
}

// runCommand executes a single linux command.
func runCommand(ctx context.Context, cli Client, container *Container, cmd []string) error {

	// Creates the instance of the process to run.
	/*exec, err := cli.ContainerExecCreate(ctx, container.ID, types.ExecConfig{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
		Privileged:   true,
	})
	if err != nil {
		//log.Print("error", "Error creating exec instance", err)
		return err
	}

	// Starts created exec instance and attaches stdout and stderr.
	response, err := cli.ContainerExecAttach(ctx, exec.ID, types.ExecConfig{})
	if err != nil {
		//log.Print("error", "Error starting exec instance", err)
		return err
	}

	// We use cmd[2] to get the actual user input command since we're running
	// all commands under bash as /bin/bash -c <command>.
	//logs.Cmd = cmd[2]
	//logs.CmdID = exec.ID
	//logs.First = true

	//log.Print("info", "Exec info", exec)

	storeLogs(ctx, cli, &response)*/

	return nil
}

// Keeps reading the stdout and stderr output of a Docker exec instance until
// the process finishes.
func storeLogs(ctx context.Context, cli Client, response *types.HijackedResponse) {

	/*proc, _ := cli.ContainerExecInspect(ctx, logs.CmdID)

	// While we haven't read all the output or the process is still running, keep
	// getting and storing output into the logs.
	for proc.Running {

		buf := make([]byte, bufferSize)

		n, _ := response.Reader.Read(buf)
		sendLogs(logs, buf[8:])

		for n == bufferSize {
			n, _ = response.Reader.Read(buf)

			// The first output of the command was already sent. Now, we just need to
			// append the subsequent outputs.
			logs.First = false

			if n > 0 {
				sendLogs(logs, buf)
			}

		}

		proc, _ = cli.ContainerExecInspect(ctx, logs.CmdID)
		// If there was no enough bytes to fill the buffer, but the process is still
		// running, wait for a couple of seconds since it might be a heavy process
		// that is not outputing anything.
		if proc.Running {
			time.Sleep(2 * time.Second)
		}

	}*/

	//log.Print("info", "Process info", proc.ExitCode, proc.Running)
}

// Stores the stdout and stderr outputs into the log database.
// We use cmd[2] to get the actual user input command since we're running
// all commands under bash as /bin/bash -c <command>.
func sendLogs(buf []byte) error {
	return nil
	//return logs.AppendCommandOutput(buildID, taskUUID, execID, buf)
}
