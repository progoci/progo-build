package build

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"
	"github.com/progoci/progo-core/uuid"

	"github.com/progoci/progo-build/pkg/docker"
)

// Build is the build manager that creates docker networks and containers.
type Build struct {
	DockerClient docker.Docker
}

// Opts is the configuration to crate a new build.
type Opts struct {
	Image string

	// Used to generate the virtual host used by nginx-proxy for reverse proxy.
	VirtualHostSuffix string

	BuildID        string
	NetworkPreffix string
}

// Response is the response to a setup.
type Response struct {
	ID           string
	ContainerIDs []string
	VirtualHost  string // The virtualhost where the live site is accessible from (used by nginx-proxy).
}

// New initilizes a new build manager.
func New(docker docker.Docker) *Build {
	return &Build{
		DockerClient: docker,
	}
}

// Setup performs all the docker-related tasks to fire up a new build.
//
// It creates a Docker network for the build, adds the networ
func (b *Build) Setup(ctx context.Context, opts *Opts) (*Response, error) {

	if _, ok := availableImages[opts.Image]; !ok {
		return nil, errors.New("image is not valid")
	}

	// Networking.
	networkingConfig, err := b.DockerClient.NetworkConnect(ctx, opts.NetworkPreffix, opts.BuildID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to network")
	}

	// Container configuration.
	containerID := uuid.Get()
	virtualhost := fmt.Sprintf("%s.%s", containerID, opts.VirtualHostSuffix)
	config := b.DockerClient.ContainerConfig(opts.Image, virtualhost)

	createOpts := &docker.CreateOpts{
		Config:           config,
		NetworkingConfig: networkingConfig,
	}
	containerID, err = b.DockerClient.ContainerCreate(ctx, createOpts)
	if err != nil {
		return nil, fmt.Errorf("could not create container: %w", err)
	}

	err = b.DockerClient.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return nil, fmt.Errorf("could not start container %s: %w", containerID, err)
	}

	return &Response{
		ID:           opts.BuildID,
		VirtualHost:  virtualhost,
		ContainerIDs: []string{containerID},
	}, nil
}
