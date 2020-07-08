package build_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"

	"github.com/progoci/progo-build/internal/types"
	"github.com/progoci/progo-build/pkg/build"
	"github.com/progoci/progo-build/pkg/docker"
	"github.com/progoci/progo-build/pkg/plugin"
)

type testPlugin struct {
	serviceNames []string
	commands     []string
}

func (t *testPlugin) Run(ctx context.Context, opts *plugin.RunOpts) {
	t.commands = append(t.commands, opts.Step.Commands...)
	t.serviceNames = append(t.serviceNames, opts.ServiceName)
}

var containers = []string{
	"testcontainer001",
	"testcontainer002",
	"testcontainer003",
}

func initDocker(cli *docker.DockerMock) *docker.DockerMock {
	cli.GenerateContainerNameFunc = func() string {
		index := len(cli.GenerateContainerNameCalls()) - 1

		return containers[index]
	}
	cli.ContainerCreateFunc = func(ctx context.Context, opts *docker.CreateOpts) (string, error) {
		index := len(cli.GenerateContainerNameCalls()) - 1

		return containers[index], nil
	}
	cli.ContainerConfigFunc = func(image string, virtualhost string) *container.Config {
		return &container.Config{}
	}
	cli.ContainerStartFunc = func(ctx context.Context, containerID string, options dockerTypes.ContainerStartOptions) error {
		return nil
	}

	cli.NetworkConnectFunc = func(ctx context.Context, preffix string, name string) (*network.NetworkingConfig, error) {
		return &network.NetworkingConfig{}, nil
	}
	cli.NetworkingConfigFunc = func(networkID string) *network.NetworkingConfig {
		return &network.NetworkingConfig{}
	}

	return cli
}

func TestSetup(t *testing.T) {

	tests := []struct {
		name              string
		services          []*types.Service
		virtualHostSuffix string
		buildID           string
		networkPreffix    string
		shouldErr         bool
		expected          *build.Build
	}{
		// Running a single service.
		{
			name: "single service using bash commands",
			services: []*types.Service{
				{
					Name:  "php",
					Image: "progoci/ubuntu18.04-php7.2-apache",
					Steps: []*types.Step{
						{
							Name:     "running command",
							Plugin:   "test",
							Commands: []string{"cat /var/log/bootstrap.log"},
						},
						{
							Name:     "running command 2",
							Plugin:   "test",
							Commands: []string{"tail -f /var/log/bootstrap.log", "echo hello"},
						},
					},
				},
			},
			virtualHostSuffix: "test.progo.ci",
			buildID:           "buildtest123456",
			networkPreffix:    "progo",
			expected: &build.Build{
				ID: "buildtest123456",
				Containers: []*build.Container{
					{ID: containers[0], VirtualHost: fmt.Sprintf("%s.%s", containers[0], "test.progo.ci")},
				},
			},
		},
		// Running two services.
		{
			name: "two services using bash commands",
			services: []*types.Service{
				{
					Name:  "php",
					Image: "progoci/ubuntu18.04-php7.2-apache",
					Steps: []*types.Step{
						{
							Name:     "running command",
							Plugin:   "test",
							Commands: []string{"cat /var/log/bootstrap.log"},
						},
					},
				},
				{
					Name:  "mysql",
					Image: "progoci/ubuntu18.04-php7.2-apache",
					Steps: []*types.Step{
						{
							Name:     "running command",
							Plugin:   "test",
							Commands: []string{"tail -f /var/log/appache.log"},
						},
					},
				},
			},
			virtualHostSuffix: "test.progo.ci",
			buildID:           "buildtest123456",
			networkPreffix:    "progo",
			expected: &build.Build{
				ID: "buildtest123456",
				Containers: []*build.Container{
					{ID: containers[0], VirtualHost: fmt.Sprintf("%s.%s", containers[0], "test.progo.ci")},
					{ID: containers[1], VirtualHost: fmt.Sprintf("%s.%s", containers[1], "test.progo.ci")},
				},
			},
		},
	}

	for _, test := range tests {

		// Docker.
		docker := &docker.DockerMock{}
		docker = initDocker(docker)

		// Test Plugin.
		pluginManager := &plugin.Manager{
			Plugins: make(map[string]plugin.Plugin),
		}
		pluginTest := &testPlugin{}
		pluginManager.Add("test", pluginTest)

		// Builder object to test.
		builder := build.New(docker, pluginManager)

		t.Run(test.name, func(t *testing.T) {

			opts := &build.Opts{
				Services:          test.services,
				VirtualHostSuffix: test.virtualHostSuffix,
				BuildID:           test.buildID,
				NetworkPreffix:    test.networkPreffix,
			}
			actual, err := builder.Setup(context.Background(), opts)

			if test.shouldErr {
				assert.Nil(t, err)
			} else {
				assert.True(t, cmp.Equal(test.expected, actual))

				// Checks the plugin was invoked.
				// TODO: Make it independent of sleep.
				time.Sleep(100 * time.Millisecond)
				var services []string
				var allCommands []string
				for _, s := range test.services {

					for _, step := range s.Steps {
						// For each step, the service name is also appended.
						services = append(services, s.Name)
						allCommands = append(allCommands, step.Commands...)
					}

				}

				assert.True(t, cmp.Equal(services, pluginTest.serviceNames))
				assert.True(t, cmp.Equal(allCommands, pluginTest.commands))
			}

		})

	}

}
