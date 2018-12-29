package main

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"

	"github.com/teamslizco/recorder/internal/allsoda"
)

type Specification struct {
	SodaEndpoint string `required:"true" split_words:"true"`
}

func main() {
	fmt.Println("Hello, recorder ;)")

	logger := logrus.New()

	var s Specification
	err := envconfig.Process("Recorder", &s)
	if err != nil {
		logger.Fatal(err.Error())
	}

	fmt.Printf("Initializing client with %s\n", s.SodaEndpoint)
	client := allsoda.New(s.SodaEndpoint, logger)

	inspecs, err := client.RetrieveInspections()
	if err != nil {
		logger.Error(err.Error())
	}

	fmt.Printf("%d inspections retrieved\n", len(inspecs.Inspections))
}
