package service

import (
	"flag"
	"net/http"
	"os"

	"progo/build/pkg/endpoint"
	servicehttp "progo/build/pkg/http"
	"progo/build/pkg/middleware"
	"progo/build/pkg/service"

	"github.com/go-kit/kit/log"
)

// Run starts the build service.
func Run(port string) {
	var (
		httpAddr = flag.String("http.addr", port, "HTTP listen address")
	)
	flag.Parse()

	logger := log.NewLogfmtLogger(os.Stderr)

	var svc service.Build
	{
		svc = service.NewBuildService()
		svc = middleware.NewLoggingMiddleware(logger, svc)
	}

	var h http.Handler
	{
		endpoints := endpoint.MakeEndpoints(svc)
		h = servicehttp.NewService(endpoints, logger)
	}

	logger.Log("msg", "Running service", "addr", port)

	server := &http.Server{
		Addr:    *httpAddr,
		Handler: h,
	}

	logger.Log("err", server.ListenAndServe())
}
