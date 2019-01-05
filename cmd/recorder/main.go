package main

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/teamslizco/recorder/internal/soda"
)

type Specification struct {
	SodaEndpoint string `required:"true" split_words:"true"`
}

func main() {

	logger := logrus.New()
	logger.Info("Hello, recorder ;)")

	var s Specification
	err := envconfig.Process("Recorder", &s)
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Infof("Initializing client with %s\n", s.SodaEndpoint)
	svc, err := soda.New(s.SodaEndpoint, logger)
	if err != nil {
		logger.Fatal(err.Error())
	}
	srv := server(svc)

	logger.Fatal(srv.ListenAndServe())
}
