package models

import (
	"strconv"
)

// Transport represents SQL table and JSON object.
type Transport struct {
	ID                  int    `json:"id" sql:"id,pk"`
	Person              string `json:"-" sql:"person"`
	RegistrationAddress string `json:"registration_address" sql:"registration_address"`
	RegistrationCode    int    `json:"registration_code" sql:"registration_code"`
	Registration        string `json:"registration" sql:"registration"`
	Date                string `json:"date" sql:"date"`
	DepCode             int    `json:"-" sql:"dep_code"`
	Dep                 string `json:"-" sql:"dep"`
	Brand               string `json:"model" sql:"brand"`
	Model               string `json:"-" sql:"model"`
	Year                int    `json:"year" sql:"year"`
	Color               string `json:"color" sql:"color"`
	Kind                string `json:"kind" sql:"kind"`
	Body                string `json:"body" sql:"body"`
	Purpose             string `json:"-" sql:"purpose"`
	Fuel                string `json:"fuel" sql:"fuel"`
	Capacity            int    `json:"capacity" sql:"capacity"`
	OwnWeight           int    `json:"own_weight" sql:"own_weight"`
	TotalWeight         int    `json:"-" sql:"total_weight"`
	Number              string `json:"number" sql:"number,notnull"`
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
