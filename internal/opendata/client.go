package opendata

import (
	"encoding/json"

	"github.com/go-resty/resty"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Client struct {
	Endpoint *string
	Logger   logrus.FieldLogger
}

func New(endpoint *string, logger logrus.FieldLogger) *Client {
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
	RecordDate           *string `json:"record_date,omitempty"`
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
	if client.Endpoint == nil {
		return nil, errors.New("No endpoint passed to retrieve inspections")
	}

	resp, err := resty.R().Get(*client.Endpoint)
	if err != nil {
		return nil, errors.Wrap(err, "Could not retrieve inspections")
	}

	if resp.Body() == nil {
		return nil, errors.New("Expected non-nil response body when retrieving inspections")
	}

	inspecs := []*Inspection{}
	err = json.Unmarshal(resp.Body(), &inspecs)
	if err != nil {
		return nil, errors.Wrap(err, "Error unmarshaling json response when retrieving inspections")
	}

	return &Output{inspecs}, err
}
