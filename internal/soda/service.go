package soda

import (
	"encoding/json"

	"github.com/SebastiaanKlippert/go-soda"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type service struct {
	Endpoint    string
	Logger      logrus.FieldLogger
	inspections []*Inspection
}

func New(endpoint string, logger logrus.FieldLogger) (*service, error) {
	svc := &service{
		Endpoint: endpoint,
		Logger:   logger,
	}
	if err := svc.LoadInspections(); err != nil {
		return nil, errors.Wrap(err, "failed to load inspections into service")
	}

	return svc, nil
}

type Inspection struct {
	CuisineDescription   string `json:"cuisine_description,omitempty"`
	DBA                  string `json:"dba,omitempty"`
	Boro                 string `json:"boro,omitempty"`
	InspectionDate       string `json:"inspection_date,omitempty"`
	Building             string `json:"building,omitempty"`
	ZipCode              string `json:"zipcode,omitempty"`
	Score                string `json:"score,omitempty"`
	Phone                string `json:"phone,omitempty"`
	Street               string `json:"street,omitempty"`
	Grade                string `json:"grade,omitempty"`
	CriticalFlag         string `json:"critical_flag,omitempty"`
	Camis                string `json:"camis,omitempty"`
	Action               string `json:"action,omitempty"`
	ViolationCode        string `json:"violation_code,omitempty"`
	ViolationDescription string `json:"violation_description,omitempty"`
	GradeDate            string `json:"grade_date,omitempty"`
	InspectionType       string `json:"inspection_type,omitempty"`
}

func (service *service) LoadInspections() error {
	if service.Endpoint == "" {
		return errors.New("No endpoint passed to retrieve inspections")
	}

	inspecs := []*Inspection{}
	gr := soda.NewGetRequest(service.Endpoint, "")
	gr.Format = "json"
	gr.Query.AddOrder("camis", false)
	ogr, err := soda.NewOffsetGetRequest(gr)
	if err != nil {
		return errors.Wrap(err, "could not initialize soda request")
	}

	for i := 0; i < 4; i++ {

		ogr.Add(1)
		go func() {
			defer ogr.Done()

			for {
				resp, err := ogr.Next(2000)
				if err == soda.ErrDone {
					break
				}
				if err != nil {
					service.Logger.Fatal(errors.Wrap(err, "Error executing soda request"))
				}

				results := []*Inspection{}
				err = json.NewDecoder(resp.Body).Decode(&results)
				resp.Body.Close()
				if err != nil {
					service.Logger.Fatal(errors.Wrap(err, "Error decoding retrieve inspections response"))
				}

				for _, insp := range results {
					inspecs = append(inspecs, insp)
				}
			}
		}()
	}
	ogr.Wait()

	service.inspections = inspecs
	return nil
}

type InspectionsInput struct {
	Limit int `json:"limit"`
}

func (service *service) Inspections(input *InspectionsInput) []*Inspection {
	if input.Limit == 0 {
		input.Limit = 10
	}
	return service.inspections[0:input.Limit]
}
