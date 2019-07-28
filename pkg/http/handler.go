package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"progo/build/pkg/endpoint"
	"progo/build/pkg/service"
)

// NewService wires endpoints to the HTTP transport.
func NewService(svcEndpoints endpoint.Endpoints, logger log.Logger) http.Handler {

	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r.Methods("POST").Path("/build").Handler(kithttp.NewServer(
		svcEndpoints.Create,
		decodeCreateRequest,
		encodeResponse,
		options...,
	))

	return r
}

type errorer interface {
	error() error
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

	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case service.ErrBuildNotFound:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
