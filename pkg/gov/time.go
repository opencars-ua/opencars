package gov

import (
	"fmt"
	"strings"
	"time"
)

type Time struct {
	time.Time
}

const TimeFormat = "2006-01-02T15:04:05.999999"

func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	if s == "null" {
		ct.Time = time.Time{}
		return
	}

	ct.Time, err = time.Parse(TimeFormat, s)

	return
}

func (ct *Time) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(TimeFormat))), nil
}
