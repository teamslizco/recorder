// +build db
package db

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) (*db, error) {
	c := &Config{
		Host:     os.Getenv("RECORDER_DB_HOST"),
		Port:     os.Getenv("RECORDER_DB_PORT"),
		User:     os.Getenv("RECORDER_DB_USER"),
		Password: os.Getenv("RECORDER_DB_PASSWORD"),
		Name:     os.Getenv("RECORDER_DB_HOST"),
		Logger:   testLogger(),
	}

	fmt.Println(c.Host)
	fmt.Println(c.Port)

	d, err := New(c)
	if err != nil {
		assert.FailNow(t, "failed to connect to test db")
	}

	return d.(*db), nil
}

func testLogger() logrus.FieldLogger {
	log := logrus.New()
	log.Out = ioutil.Discard
	return log
}

func (d *db) seedCreateInspectionData(t *testing.T) {
	grade := &Grade{Name: "TEST-GRADE"}
	boro := &Boro{Name: "TEST-BORO"}

	d.db.Create(grade)
	d.db.Create(boro)
}

func TestCreateInspection(t *testing.T) {
	data, err := setupTestDB(t)
	if err != nil {
		assert.FailNow(t, "failed to start test db")
	}

	data.seedCreateInspectionData(t)

	inspec := &CreateInspectionInput{
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

	resp, err := data.CreateInspection(inspec)

	assert.Nil(t, err)
	if !assert.NotNil(t, resp) {
		assert.FailNow(t, "expected create inspection response to be non-nil")
	}
	assert.NotZero(t, resp.ID)
}
