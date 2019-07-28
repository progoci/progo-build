package service

import (
	"context"
	"errors"

	"progo/build/pkg/entity"
)

var (
	// ErrBuildNotFound is thrown when build was not found.
	ErrBuildNotFound = errors.New("build not found")
)

// Build describes the build service.
type Build interface {
	Create(ctx context.Context, build entity.Build) (string, error)
}

type buildService struct{}

// NewBuild creates a new build service.
func NewBuild() Build {
	return &buildService{}
}

// Create stores a new build entity.
func (c *buildService) Create(ctx context.Context,
	container entity.Build) (string, error) {

	return "id", nil
}
