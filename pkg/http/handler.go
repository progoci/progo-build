package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"progo/build/pkg/endpoint"
)

// NewService wires endpoints to the HTTP transport.
func NewService(svcEndpoints endpoint.Endpoints, logger log.Logger) http.Handler {

	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
	}

	createHandler := kithttp.NewServer(
		svcEndpoints.Create,
		decodeCreateRequest,
		encodeResponse,
		options...,
	)

	r.Methods("POST").Path("/build").Handler(createHandler)

	return r
}

func decodeCreateRequest(_ context.Context,
	r *http.Request) (request interface{}, err error) {

	var req endpoint.CreateRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Build); e != nil {
		return nil, e
	}

	return req, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter,
	response interface{}) error {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return json.NewEncoder(w).Encode(response)
}
