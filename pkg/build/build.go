package build

import (
	"context"
	"fmt"

	dockerTypes "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/pkg/errors"

	"github.com/progoci/progo-build/internal/types"
	"github.com/progoci/progo-build/pkg/docker"
	"github.com/progoci/progo-build/pkg/plugin"
	"github.com/progoci/progo-core/uuid"
)

// Manager is the build manager that creates docker networks and containers.
type Manager struct {
	DockerClient  docker.Docker
	PluginManager *plugin.Manager
}

// Opts is the configuration to crate a new build.
type Opts struct {
	Services []*types.Service

	// Used to generate the virtual host used by nginx-proxy for reverse proxy.
	VirtualHostSuffix string

	BuildID        string
	NetworkPreffix string
}

// Container is information about the containers in a build.
type Container struct {
	ID          string
	VirtualHost string
}

// Build is the response to a setup.
type Build struct {
	ID         string
	Containers []*Container
}

// New initilizes a new build manager.
func New(docker docker.Docker, pluginManager *plugin.Manager) *Manager {
	return &Manager{
		DockerClient:  docker,
		PluginManager: pluginManager,
	}
}

// Setup performs all the docker-related tasks to fire up a new build.
//
// It creates a Docker network for the build to avoid communication between
// containers from different builds. It then creates the containers for the
// build and starts them.
func (m *Manager) Setup(ctx context.Context, opts *Opts) (*Build, error) {

	invalidImage := invalidImage(opts.Services)
	if invalidImage != "" {
		msg := fmt.Sprintf("image %s is not valid", invalidImage)
		return nil, errors.New(msg)
	}

	// Networking.
	networkingConfig, err := m.DockerClient.NetworkConnect(ctx, opts.NetworkPreffix, opts.BuildID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to network")
	}

	response := &Build{
		ID:         opts.BuildID,
		Containers: []*Container{},
	}

	for _, service := range opts.Services {
		container, err := m.setupContainer(ctx, service.Image, opts.VirtualHostSuffix, networkingConfig)
		if err != nil {
			return nil, errors.Wrapf(err, "could not setup service %s", service.Name)
		}

		go m.ExecuteSteps(ctx, opts.BuildID, service.Name, container.ID, service.Steps)

		response.Containers = append(response.Containers, container)
	}

	return response, nil
}

// setupService setups a single service.
func (m *Manager) setupContainer(ctx context.Context, image string,
	virtualhostSuffix string, networking *network.NetworkingConfig) (*Container, error) {

	// Container configuration.
	containerID := uuid.Get()
	virtualhost := fmt.Sprintf("%s.%s", containerID, virtualhostSuffix)
	config := m.DockerClient.ContainerConfig(image, virtualhost)

	createOpts := &docker.CreateOpts{
		Config:           config,
		NetworkingConfig: networking,
	}
	containerID, err := m.DockerClient.ContainerCreate(ctx, createOpts)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create container %s", containerID)
	}

	err = m.DockerClient.ContainerStart(ctx, containerID, dockerTypes.ContainerStartOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not start container %s: %w", containerID, err)
	}

	return &Container{
		ID:          containerID,
		VirtualHost: virtualhost,
	}, nil
}

// ExecuteSteps runs the steps in the configuration for a single service.
func (m *Manager) ExecuteSteps(ctx context.Context, buildID string,
	serviceName string, containerID string, steps []*types.Step) {

	for i, step := range steps {
		m.PluginManager.Run(ctx, buildID, serviceName, containerID, i+1, step)
	}

}
