package models

import (
	"strconv"
)

// Transport represents SQL table and JSON object.
type Transport struct {
	ID                  int    `json:"id" db:"id,pk"`
	Person              string `json:"-" db:"person"`
	RegistrationAddress string `json:"registration_address" db:"registration_address"`
	RegistrationCode    int    `json:"registration_code" db:"registration_code"`
	Registration        string `json:"registration" db:"registration"`
	Date                string `json:"date" db:"date"`
	DepCode             int    `json:"-" db:"dep_code"`
	Dep                 string `json:"-" db:"dep"`
	Brand               string `json:"model" db:"brand"`
	Model               string `json:"-" db:"model"`
	Year                int    `json:"year" db:"year"`
	Color               string `json:"color" db:"color"`
	Kind                string `json:"kind" db:"kind"`
	Body                string `json:"body" db:"body"`
	Purpose             string `json:"-" db:"purpose"`
	Fuel                string `json:"fuel" db:"fuel"`
	Capacity            int    `json:"capacity" db:"capacity"`
	OwnWeight           int    `json:"own_weight" db:"own_weight"`
	TotalWeight         int    `json:"-" db:"total_weight"`
	Number              string `json:"number" db:"number,notnull"`
}

// Valid checks whatever transport number valid or not.
func (transport Transport) Valid() (matched bool) {
	return transport.Number != "NULL"
}

// TrimNull returns empty string in case of NULL.
func TrimNull(s string) string {
	if s == "NULL" {
		return ""
	}

	return s
}

// NewTransportFromCSV parses CSV line into transport structure.
func NewTransportFromCSV(record []string) *Transport {
	transport := new(Transport)

	transport.Person = record[0]
	transport.RegistrationAddress = record[1]
	transport.RegistrationCode, _ = strconv.Atoi(record[2])
	transport.Registration = record[3]
	transport.Date = record[4]
	transport.DepCode, _ = strconv.Atoi(record[5])
	transport.Dep = TrimNull(record[6])
	transport.Brand = TrimNull(record[7])
	transport.Model = TrimNull(record[8])
	transport.Year, _ = strconv.Atoi(record[9])
	transport.Color = TrimNull(record[9])
	transport.Kind = TrimNull(record[11])
	transport.Body = TrimNull(record[12])
	transport.Purpose = TrimNull(record[13])
	transport.Fuel = TrimNull(record[14])
	transport.Capacity, _ = strconv.Atoi(record[15])
	transport.OwnWeight, _ = strconv.Atoi(record[16])
	transport.TotalWeight, _ = strconv.Atoi(record[17])
	transport.Number = TrimNull(record[18])

	return transport
}
