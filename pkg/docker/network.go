package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/pkg/errors"
)

// NetworkConnect handles networking of the containers.
//
// A build might have several containers, so this method creates a network for
// to allow communication among them. It also connects the nginx proxy to the
// new for auto-proxying.
//
// For instance, the preffix can be the build id and the name the container id.
func (client *Client) NetworkConnect(ctx context.Context, preffix string, name string) (*network.NetworkingConfig, error) {
	networkID := fmt.Sprintf("%s_%s", preffix, name)
	_, err := client.Conn.NetworkCreate(ctx, networkID, types.NetworkCreate{})
	if err != nil {
		return nil, errors.Wrap(err, "could not create network")
	}

	// Connect proxy to new network.
	err = client.Conn.NetworkConnect(ctx, networkID, client.ProxyContainerID, &network.EndpointSettings{})
	if err != nil {
		return nil, errors.Wrap(err, "could not connect proxy to new network")
	}

	return client.NetworkingConfig(networkID), nil
}

// NetworkingConfig generates a network configuration for a container.
func (client *Client) NetworkingConfig(networkID string) *network.NetworkingConfig {
	return &network.NetworkingConfig{
		EndpointsConfig: map[string]*network.EndpointSettings{
			"net": &network.EndpointSettings{
				NetworkID: networkID,
			},
		},
	}
}
