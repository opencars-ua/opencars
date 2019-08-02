package model

import (
	"strconv"
	"strings"
)

// Operation represents public registrations of transport.
type Operation struct {
	ID          int    `json:"-" sql:"id,pk"`
	Person      string `json:"person" sql:"person"`
	Address     string `json:"address" sql:"reg_addr_koatuu"`
	Code        int    `json:"operation" sql:"code"`
	Description string `json:"description" sql:"description"`
	Date        string `json:"date" sql:"date"`
	OfficeID    int    `json:"office_id" sql:"office_id"`
	OfficeName  string `json:"office_name" sql:"office_name"`
	Brand       string `json:"brand" sql:"brand" `
	Model       string `json:"model" sql:"model"`
	Year        int    `json:"year" sql:"year"`
	Color       string `json:"color" sql:"color"`
	Kind        string `json:"kind" sql:"kind"`
	Body        string `json:"body" sql:"body"`
	Purpose     string `json:"purpose" sql:"purpose"`
	Fuel        string `json:"fuel" sql:"fuel"`
	Capacity    int    `json:"capacity" sql:"capacity"`
	Weight      int    `json:"weight" sql:"weight"`
	Number      string `json:"number" sql:"number,notnull"`
	VIN         string `json:"vin,omitempty" sql:"vin"`
}

// Valid checks whatever model number valid or not.
func (op Operation) Valid() (matched bool) {
	return op.Number != ""
}

// TrimNull returns empty string in case of NULL.
func TrimNull(s string) string {
	s = strings.TrimSpace(s)

	if s == "NULL" {
		return ""
	}

	return s
}

func (op *Operation) fixBrand() {
	op.Brand = strings.TrimSpace(op.Brand)
	op.Brand = strings.TrimSuffix(op.Brand, op.Model)
	op.Brand = strings.TrimSpace(op.Brand)
}

func (op *Operation) fixDescription() {
	op.Description = strings.TrimSpace(op.Description)
	op.Description = strings.TrimPrefix(op.Description, strconv.Itoa(op.Code))
	op.Description = strings.TrimSpace(op.Description)
	op.Description = strings.TrimPrefix(op.Description, "-")
	op.Description = strings.TrimSpace(op.Description)
}

// NewOperation parses CSV line into operation model.
func NewOperation(record []string) *Operation {
	o := new(Operation)

	o.Person = record[0]
	o.Address = strings.TrimSpace(record[1])
	o.Code, _ = strconv.Atoi(record[2])
	o.Description = strings.ToUpper(strings.TrimSpace(record[3]))
	o.Date = strings.TrimSpace(record[4])
	o.OfficeID, _ = strconv.Atoi(record[5])
	o.OfficeName = TrimNull(record[6])
	o.Brand = strings.ToUpper(TrimNull(record[7]))
	o.Model = strings.ToUpper(TrimNull(record[8]))
	o.Year, _ = strconv.Atoi(record[9])
	o.Color = strings.ToUpper(TrimNull(record[10]))
	o.Kind = strings.ToUpper(TrimNull(record[11]))
	o.Body = strings.ToUpper(TrimNull(record[12]))
	o.Purpose = strings.ToUpper(TrimNull(record[13]))
	o.Fuel = strings.ToUpper(TrimNull(record[14]))
	o.Capacity, _ = strconv.Atoi(record[15])
	o.Weight, _ = strconv.Atoi(record[16])
	o.Number = TrimNull(record[18])

	o.fixBrand()
	o.fixDescription()

	return o
}
