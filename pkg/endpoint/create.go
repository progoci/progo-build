package endpoint

import (
	"context"

	endpoint "github.com/go-kit/kit/endpoint"

	"progo/build/pkg/entity"
	"progo/build/pkg/service"
)

// CreateRequest collects the request parameters for the Create method.
type CreateRequest struct {
	Build entity.Build
}

// CreateResponse collects the response parameters for the Create method.
type CreateResponse struct {
	ID  string `json:"id"`
	Err error  `json:"error,omitempty"`
}

func makeCreateEndpoint(s service.Build) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := s.Create(ctx, req.Build)

		return CreateResponse{ID: id, Err: err}, nil
	}
}
