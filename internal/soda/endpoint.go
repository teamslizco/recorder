package soda

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type Service interface {
	Inspections(*InspectionsInput) []*Inspection
}

type inspectionsResponse struct {
	Inspections []*Inspection `json:"inspections"`
	Err         string        `json:"error,omitempty"`
}

func MakeInspectionsEndpoint(svc Service) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		input := request.(InspectionsInput)
		v := svc.Inspections(&input)
		return inspectionsResponse{v, ""}, nil
	}
}

func DecodeInspectionsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req InspectionsInput
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
