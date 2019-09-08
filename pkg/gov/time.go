package gov

import (
	"fmt"
	"strings"
	"time"
)

// Time is an extension over time.Time{}.
// Created to override JSON serialization methods.
type Time struct{ time.Time }

const timeFormat = "2006-01-02T15:04:05.999999"

// UnmarshalJSON deserialize incoming time into custom time format.
func (ct *Time) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")

	if s == "null" {
		ct.Time = time.Time{}
		return
	}

	ct.Time, err = time.Parse(timeFormat, s)

	return
}

// MarshalJSON serializes time.
func (ct *Time) MarshalJSON() ([]byte, error) {
	if ct.Time.UnixNano() == (time.Time{}).UnixNano() {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(timeFormat))), nil
}
