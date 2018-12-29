package db

import (
	"strconv"
	"time"

	"github.com/pkg/errors"
)

func toStrPtr(str string) *string {
	return &str
}

func strToDate(str *string) (time.Time, error) {
	var t time.Time
	if str == nil {
		return t, errors.New("cannot convert nil-value to Time")
	}

	t, err := time.Parse(time.RFC3339, *str)
	if err != nil {
		return t, errors.Wrap(err, "encountered error converting string to date")
	}

	return t, nil
}

func stringP(v *string) string {
	if v == nil {
		return ""
	}
	return *v
}

func atoi(v *string) int {
	if v == nil {
		return 0
	}

	i, _ := strconv.Atoi(*v)
	return i
}
