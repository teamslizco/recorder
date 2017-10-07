package opendata

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRetrieveInspections_shouldRetrieveAndUnmarshalResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body := `[
			      {
			      "cuisine_description" : "Bakery",
			      "dba" : "MORRIS PARK BAKE SHOP",
			      "record_date" : "2017-09-15T06:00:47",
			      "boro" : "BRONX",
			      "inspection_date" : "2017-05-18T00:00:00",
			      "building" : "1007",
			      "zipcode" : "10462",
			      "score" : "7",
			      "phone" : "7188924968",
			      "street" : "MORRIS PARK AVE",
			      "grade" : "A",
			      "critical_flag" : "Critical",
			      "camis" : "30075445",
			      "action" : "Violations were cited in the following area(s).",
			      "violation_code" : "06D",
			      "violation_description" : "Food contact surface not properly washed, rinsed and sanitized after each use and following any activity when contamination may have occurred.",
			      "grade_date" : "2017-05-18T00:00:00",
			      "inspection_type" : "Cycle Inspection / Initial Inspection"
			      }
			      , {
			      "cuisine_description" : "Bakery",
			      "dba" : "MORRIS PARK BAKE SHOP",
			      "record_date" : "2017-09-15T06:00:47",
			      "boro" : "BRONX",
			      "inspection_date" : "2017-05-18T00:00:00",
			      "building" : "1007",
			      "zipcode" : "10462",
			      "score" : "7",
			      "phone" : "7188924968",
			      "street" : "MORRIS PARK AVE",
			      "grade" : "A",
			      "critical_flag" : "Not Critical",
			      "camis" : "30075445",
			      "action" : "Violations were cited in the following area(s).",
			      "violation_code" : "10F",
			      "violation_description" : "Non-food contact surface improperly constructed. Unacceptable material used. Non-food contact surface or equipment improperly maintained and/or not properly sealed, raised, spaced or movable to allow accessibility for cleaning on all sides, above and underneath the unit.",
			      "grade_date" : "2017-05-18T00:00:00",
			      "inspection_type" : "Cycle Inspection / Initial Inspection"
			      }
			]`
		fmt.Fprintln(w, body)
	}))

	defer ts.Close()

	expected := &Output{
		Inspections: []*Inspection{

			{
				CuisineDescription:   toStrPtr("Bakery"),
				DBA:                  toStrPtr("MORRIS PARK BAKE SHOP"),
				RecordDate:           toStrPtr("2017-09-15T06:00:47"),
				Boro:                 toStrPtr("BRONX"),
				InspectionDate:       toStrPtr("2017-05-18T00:00:00"),
				Building:             toStrPtr("1007"),
				ZipCode:              toStrPtr("10462"),
				Score:                toStrPtr("7"),
				Phone:                toStrPtr("7188924968"),
				Street:               toStrPtr("MORRIS PARK AVE"),
				Grade:                toStrPtr("A"),
				CriticalFlag:         toStrPtr("Critical"),
				Camis:                toStrPtr("30075445"),
				Action:               toStrPtr("Violations were cited in the following area(s)."),
				ViolationCode:        toStrPtr("06D"),
				ViolationDescription: toStrPtr("Food contact surface not properly washed, rinsed and sanitized after each use and following any activity when contamination may have occurred."),
				GradeDate:            toStrPtr("2017-05-18T00:00:00"),
				InspectionType:       toStrPtr("Cycle Inspection / Initial Inspection"),
			},

			{
				CuisineDescription:   toStrPtr("Bakery"),
				DBA:                  toStrPtr("MORRIS PARK BAKE SHOP"),
				RecordDate:           toStrPtr("2017-09-15T06:00:47"),
				Boro:                 toStrPtr("BRONX"),
				InspectionDate:       toStrPtr("2017-05-18T00:00:00"),
				Building:             toStrPtr("1007"),
				ZipCode:              toStrPtr("10462"),
				Score:                toStrPtr("7"),
				Phone:                toStrPtr("7188924968"),
				Street:               toStrPtr("MORRIS PARK AVE"),
				Grade:                toStrPtr("A"),
				CriticalFlag:         toStrPtr("Not Critical"),
				Camis:                toStrPtr("30075445"),
				Action:               toStrPtr("Violations were cited in the following area(s)."),
				ViolationCode:        toStrPtr("10F"),
				ViolationDescription: toStrPtr("Non-food contact surface improperly constructed. Unacceptable material used. Non-food contact surface or equipment improperly maintained and/or not properly sealed, raised, spaced or movable to allow accessibility for cleaning on all sides, above and underneath the unit."),
				GradeDate:            toStrPtr("2017-05-18T00:00:00"),
				InspectionType:       toStrPtr("Cycle Inspection / Initial Inspection"),
			},
		},
	}

	client := New(&ts.URL, nil)

	resp, err := client.RetrieveInspections()

	assert.NoError(t, err)
	assert.Equal(t, expected, resp)
}
