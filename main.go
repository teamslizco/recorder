package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/teamslizco/recorder/internal/opendata"
)

type Specification struct {
	OpenDataEndpoint *string `required:"true" split_words:"true"`
}

func main() {
	fmt.Println("Hello, recorder ;)")

	logger := logrus.New()

	var s Specification
	err := envconfig.Process("Recorder", &s)
	if err != nil {
		logger.Fatal(err.Error())
	}

	fmt.Printf("Initializing client with %s\n", s.OpenDataEndpoint)
	client := opendata.New(s.OpenDataEndpoint, logger)

	inspecs, err := client.RetrieveInspections()
	if err != nil {
		logger.Error(err.Error())
	}

	fmt.Printf("%d inspections retrieved\n", len(inspecs.Inspections))
}
