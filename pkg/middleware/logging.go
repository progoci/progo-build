package middleware

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"progo/build/pkg/entity"
	"progo/build/pkg/service"
)

type loggingMiddleware struct {
	logger log.Logger
	next   service.Build
}

func (mw loggingMiddleware) Create(ctx context.Context,
	build *entity.Build) (output string, err error) {

	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "build",
			"input", build,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Create(ctx, build)
	return
}

// NewLoggingMiddleware creates a new loggin middleware.
func NewLoggingMiddleware(logger log.Logger, svc service.Build) service.Build {
	return &loggingMiddleware{logger, svc}
}
