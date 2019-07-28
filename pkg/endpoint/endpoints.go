package endpoint

import (
	endpoint "github.com/go-kit/kit/endpoint"

	"progo/build/pkg/service"
)

// Endpoints holds all the endpoints for the build service.
type Endpoints struct {
	Create endpoint.Endpoint
}

// MakeEndpoints initializes all endpoints for the build service.
func MakeEndpoints(s service.Build) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
	}
}
