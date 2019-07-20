package service

import (
	"context"
	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(ContainerService) ContainerService

type loggingMiddleware struct {
	logger log.Logger
	next   ContainerService
}

// LoggingMiddleware takes a logger as a dependency
// and returns a ContainerService Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next ContainerService) ContainerService {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Create(ctx context.Context, s string) (e0 error) {
	defer func() {
		l.logger.Log("method", "Create", "s", s, "e0", e0)
	}()
	return l.next.Create(ctx, s)
}
