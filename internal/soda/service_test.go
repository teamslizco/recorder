package soda

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNew_shouldRetrieveAndUnmarshalResponse(t *testing.T) {
	svc, err := New("https://data.cityofnewyork.us/resource/9w7m-hzhe", logrus.New())

	assert.NoError(t, err)
	if !assert.NotNil(t, svc.Inspections) {
		assert.FailNow(t, "expected Inspections to be non-nil")
	}

	assert.True(t, len(svc.inspections) > 2000)
	t.Logf("Retrieved %d inspections in testing\n", len(svc.inspections))
}
