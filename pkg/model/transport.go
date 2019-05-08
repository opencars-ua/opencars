package model

import (
	"strconv"
	"strings"
)

// Operation represents public registrations of transport.
type Operation struct {
	ID          int    `json:"-" db:"id,pk"`
	Person      string `json:"person" db:"person"`
	Address     string `json:"address" db:"reg_addr_koatuu"`
	Code        int    `json:"operation" db:"code"`
	Description string `json:"description" db:"description"`
	Date        string `json:"date" db:"date"`
	OfficeID    int    `json:"office_id" db:"office_id"`
	OfficeName  string `json:"office_name" db:"office_name"`
	Brand       string `json:"brand" db:"brand" `
	Model       string `json:"model" db:"model"`
	Year        int    `json:"year" db:"year"`
	Color       string `json:"color" db:"color"`
	Kind        string `json:"kind" db:"kind"`
	Body        string `json:"body" db:"body"`
	Purpose     string `json:"purpose" db:"purpose"`
	Fuel        string `json:"fuel" db:"fuel"`
	Capacity    int    `json:"capacity" db:"capacity"`
	Weight      int    `json:"weight" db:"weight"`
	Number      string `json:"number" db:"number,notnull"`
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
