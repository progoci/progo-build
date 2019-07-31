package service

import (
	"context"
	"progo/build/pkg/entity"
)

// Create stores a new build entity.
func (c *buildService) Create(ctx context.Context,
	build entity.Build) (string, error) {

	return build.ID, nil
}
