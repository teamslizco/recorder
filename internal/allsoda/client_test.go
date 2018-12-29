package allsoda

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestRetrieveInspections_shouldRetrieveAndUnmarshalResponse(t *testing.T) {
	client := New(toStrPtr("https://data.cityofnewyork.us/resource/9w7m-hzhe.json"), logrus.New())

	expected := &Inspection{
		CuisineDescription:   toStrPtr("Bakery"),
		DBA:                  toStrPtr("MORRIS PARK BAKE SHOP"),
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
	}

	resp, err := client.RetrieveInspections()

	assert.NoError(t, err)
	if !assert.NotNil(t, resp.Inspections) {
		assert.FailNow(t, "expected Inspections to be non-nil")
	}

	assert.Contains(t, resp.Inspections, expected)
	assert.True(t, len(resp.Inspections) > 1000)
	fmt.Printf("Retrieved %d inspections in testing\n", len(resp.Inspections))
}
