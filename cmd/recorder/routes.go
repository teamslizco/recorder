package main

import (
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/teamslizco/recorder/internal/soda"
)

type route struct {
	path    string
	handler http.Handler
}

func routes(svc soda.Service) []route {
	return []route{
		{
			"/inspections",
			httptransport.NewServer(
				soda.MakeInspectionsEndpoint(svc),
				soda.DecodeInspectionsRequest,
				encodeResponse,
			),
		},
	}
}

func router(svc soda.Service) *mux.Router {
	rter := mux.NewRouter()
	rts := routes(svc)

	for _, r := range rts {
		rter.Handle(r.path, r.handler)
	}

	return rter
}

func server(svc soda.Service) *http.Server {
	return &http.Server{
		Handler:      router(svc),
		Addr:         "0.0.0.0:8080",
		WriteTimeout: 60 * time.Second,
		ReadTimeout:  60 * time.Second,
	}
}
