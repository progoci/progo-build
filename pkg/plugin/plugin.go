package plugin

import (
	"context"

	"github.com/pkg/errors"

	"github.com/progoci/progo-build/internal/types"
	"github.com/progoci/progo-build/pkg/docker"
)

// Plugin is the interface all plugins must implement.
type Plugin interface {
	Run(ctx context.Context, opts *RunOpts)
}

// Manager contains information about plugins.
type Manager struct {
	DockerClient docker.Docker
	Plugins      map[string]Plugin
}

// RunOpts are the options for running a step.
type RunOpts struct {
	BuildID      string
	ServiceName  string
	ContainerID  string
	StepNumber   int32
	Step         *types.Step
	DockerClient docker.Docker
}

// NewManager initializes a new Plugin Manager
func NewManager(docker docker.Docker) *Manager {
	return &Manager{
		DockerClient: docker,
		Plugins:      make(map[string]Plugin),
	}
}

// Add adds a new plugin to the list.
func (manager *Manager) Add(key string, plugin Plugin) {
	manager.Plugins[key] = plugin
}

// Get returns a plugin based on the key.
func (manager *Manager) Get(key string) Plugin {
	if plugin, ok := manager.Plugins[key]; ok {
		return plugin
	}

	return nil
}

// Run runs a step using the correct plugin.
func (manager *Manager) Run(ctx context.Context, buildID string, serviceName string,
	containerID string, stepNumber int, step *types.Step) error {

	plugin := manager.Get(step.Plugin)
	if plugin == nil {
		return errors.New("plugin does not exist")
	}

	opts := &RunOpts{
		BuildID:      buildID,
		ServiceName:  serviceName,
		ContainerID:  containerID,
		StepNumber:   int32(stepNumber),
		Step:         step,
		DockerClient: manager.DockerClient,
	}

	plugin.Run(ctx, opts)

	return nil
}
