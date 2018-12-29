package allsoda

import (
	"encoding/json"

	"github.com/SebastiaanKlippert/go-soda"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Endpoint string
	Logger   logrus.FieldLogger
}

func New(endpoint string, logger logrus.FieldLogger) *Client {
	return &Client{
		Endpoint: endpoint,
		Logger:   logger,
	}
}

type Output struct {
	Inspections []*Inspection
}

type Inspection struct {
	CuisineDescription   *string `json:"cuisine_description,omitempty"`
	DBA                  *string `json:"dba,omitempty"`
	Boro                 *string `json:"boro,omitempty"`
	InspectionDate       *string `json:"inspection_date,omitempty"`
	Building             *string `json:"building,omitempty"`
	ZipCode              *string `json:"zipcode,omitempty"`
	Score                *string `json:"score,omitempty"`
	Phone                *string `json:"phone,omitempty"`
	Street               *string `json:"street,omitempty"`
	Grade                *string `json:"grade,omitmepty"`
	CriticalFlag         *string `json:"critical_flag,omitmepty"`
	Camis                *string `json:"camis,omitmepty"`
	Action               *string `json:"action,omitmepty"`
	ViolationCode        *string `json:"violation_code,omitempty"`
	ViolationDescription *string `json:"violation_description,omitempty"`
	GradeDate            *string `json:"grade_date,omitempty"`
	InspectionType       *string `json:"inspection_type,omitempty"`
}

func (client *Client) RetrieveInspections() (*Output, error) {
	if client.Endpoint == "" {
		return nil, errors.New("No endpoint passed to retrieve inspections")
	}

	inspecs := []*Inspection{}
	gr := soda.NewGetRequest(client.Endpoint, "")
	gr.Format = "json"
	gr.Query.AddOrder("camis", false)
	ogr, err := soda.NewOffsetGetRequest(gr)
	if err != nil {
		return nil, errors.Wrap(err, "could not initialize soda request")
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
					client.Logger.Fatal(errors.Wrap(err, "Error executing soda request"))
				}

				results := []*Inspection{}
				err = json.NewDecoder(resp.Body).Decode(&results)
				resp.Body.Close()
				if err != nil {
					client.Logger.Fatal(errors.Wrap(err, "Error decoding retrieve inspections response"))
				}

				for _, insp := range results {
					inspecs = append(inspecs, insp)
				}
			}
		}()
	}
	ogr.Wait()

	return &Output{inspecs}, err
}
