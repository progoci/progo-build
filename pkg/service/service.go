package service

import "context"

// ContainerService describes the container service.
type ContainerService interface {
	Create(ctx context.Context, s string) error
}

type basicContainerService struct{}

func (b *basicContainerService) Create(ctx context.Context, s string) (e0 error) {
	// TODO implement the business logic of Create
	return e0
}

// NewBasicContainerService returns a naive, stateless implementation of ContainerService.
func NewBasicContainerService() ContainerService {
	return &basicContainerService{}
}

// New returns a ContainerService with all of the expected middleware wired in.
func New(middleware []Middleware) ContainerService {
	var svc ContainerService = NewBasicContainerService()
	for _, m := range middleware {
		svc = m(svc)
	}
	return svc
}
