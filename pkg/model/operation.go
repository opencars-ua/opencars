package model

import (
	"strconv"
	"strings"
)

// Operation represents public registrations of transport.
type Operation struct {
	ID          int64  `json:"-" sql:"id,pk"`
	Person      string `json:"person" sql:"person"`
	Address     string `json:"address" sql:"reg_addr_koatuu"`
	Code        int32  `json:"operation" sql:"code"`
	Description string `json:"description" sql:"description"`
	Date        string `json:"date" sql:"date"`
	OfficeID    int64  `json:"office_id" sql:"office_id"`
	OfficeName  string `json:"office_name" sql:"office_name"`
	Brand       string `json:"brand" sql:"brand" `
	Model       string `json:"model" sql:"model"`
	Year        int32  `json:"year" sql:"year"`
	Color       string `json:"color" sql:"color"`
	Kind        string `json:"kind" sql:"kind"`
	Body        string `json:"body" sql:"body"`
	Purpose     string `json:"purpose" sql:"purpose"`
	Fuel        string `json:"fuel" sql:"fuel"`
	Capacity    int32  `json:"capacity" sql:"capacity"`
	Weight      int32  `json:"weight" sql:"weight"`
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
	code := strconv.FormatInt(int64(op.Code), 10)

	op.Description = strings.TrimSpace(op.Description)
	op.Description = strings.TrimPrefix(op.Description, code)
	op.Description = strings.TrimSpace(op.Description)
	op.Description = strings.TrimPrefix(op.Description, "-")
	op.Description = strings.TrimSpace(op.Description)
}

// NewOperation parses CSV line into operation model.
func NewOperation(record []string) *Operation {
	o := new(Operation)

	o.Person = record[0]
	o.Address = strings.TrimSpace(record[1])

	code, _ := strconv.Atoi(record[2])
	o.Code = int32(code)

	officeID, _ := strconv.Atoi(record[5])
	o.OfficeID = int64(officeID)

	year, _ := strconv.Atoi(record[9])
	o.Year = int32(year)

	capacity, _ := strconv.Atoi(record[15])
	o.Capacity = int32(capacity)

	weight, _ := strconv.Atoi(record[16])
	o.Weight = int32(weight)

	o.Description = strings.ToUpper(strings.TrimSpace(record[3]))
	o.Date = strings.TrimSpace(record[4])
	o.OfficeName = TrimNull(record[6])
	o.Brand = strings.ToUpper(TrimNull(record[7]))
	o.Model = strings.ToUpper(TrimNull(record[8]))
	o.Color = strings.ToUpper(TrimNull(record[10]))
	o.Kind = strings.ToUpper(TrimNull(record[11]))
	o.Body = strings.ToUpper(TrimNull(record[12]))
	o.Purpose = strings.ToUpper(TrimNull(record[13]))
	o.Fuel = strings.ToUpper(TrimNull(record[14]))
	o.Number = TrimNull(record[18])

	o.fixBrand()
	o.fixDescription()

	return o
}
