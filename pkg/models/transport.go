package models

import (
	"strconv"
	"strings"
)

type TransportData struct {
	ID     int    `json:"id"`
	Number string `json:"number"`
}

type TransportResponse struct {
	Count int             `json:"count"`
	Data  []TransportData `json:"data"`
}

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

func (transport Transport) Valid() (matched bool) {
	return transport.Number != "NULL"
}

func NewTransportFromCSV(record []string) *Transport {
	transport := new(Transport)

	transport.Person = record[0]
	transport.RegistrationAddress = record[1]
	transport.RegistrationCode, _ = strconv.Atoi(record[2])
	transport.Registration = record[3]
	transport.Date = record[4]
	transport.DepCode, _ = strconv.Atoi(record[5])
	transport.Dep = strings.Trim(record[6], "NULL")
	transport.Brand = strings.Trim(record[7], "NULL")
	transport.Model = strings.Trim(record[8], "NULL")
	transport.Year, _ = strconv.Atoi(record[9])
	transport.Color = strings.Trim(record[10], "NULL")
	transport.Kind = strings.Trim(record[11], "NULL")
	transport.Body = strings.Trim(record[12], "NULL")
	transport.Purpose = strings.Trim(record[13], "NULL")
	transport.Fuel = strings.Trim(record[14], "NULL")
	transport.Capacity, _ = strconv.Atoi(record[15])
	transport.OwnWeight, _ = strconv.Atoi(record[16])
	transport.TotalWeight, _ = strconv.Atoi(record[17])
	transport.Number = strings.Trim(record[18], "NULL")

	return transport
}
