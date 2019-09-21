package model

import (
	"strconv"
	"strings"

	"github.com/opencars/opencars/pkg/hsc"
)

// Registration represents information from vehicle registration document.
type Registration struct {
	ID          int64  `json:"-"            sql:"id,pk"`
	Brand       string `json:"brand"        sql:"brand"`
	Capacity    int32  `json:"capacity"     sql:"capacity"`
	Color       string `json:"color"        sql:"color"`
	FirstReg    string `json:"first_reg"    sql:"first_reg"`
	Date        string `json:"date"         sql:"date"`
	Fuel        string `json:"fuel"         sql:"fuel"`
	Kind        string `json:"kind"         sql:"kind"`
	Body        string `json:"body"         sql:"body"`
	Year        int32  `json:"year"         sql:"year"`
	Model       string `json:"model"        sql:"model"`
	Code        string `json:"code"         sql:"code"`
	Number      string `json:"number"       sql:"number"`
	TotalWeight int32  `json:"total_weight" sql:"total_weight"`
	OwnWeight   int32  `json:"own_weight"   sql:"own_weight"`
	Category    string `json:"category"     sql:"category"`
	VIN         string `json:"vin"          sql:"vin"`
}

// RegFromHSC creates new Registration from specified hsc.Registration.
func RegFromHSC(obj *hsc.Registration) *Registration {
	r := new(Registration)

	kind := strings.Split(obj.Kind, " ")

	r.Brand = obj.Brand
	r.Color = obj.Color
	r.FirstReg = obj.DFirstReg
	r.Date = obj.DReg
	r.Fuel = obj.Fuel
	r.Kind = kind[0]
	r.Body = kind[1]
	r.Model = obj.Model
	r.Code = obj.SDoc + obj.NDoc
	r.Number = obj.NRegNew
	r.Code = obj.SDoc + obj.NDoc
	r.Category = obj.RankCategory
	r.VIN = obj.VIN

	capacity, _ := strconv.Atoi(obj.Capacity)
	r.Capacity = int32(capacity)

	year, _ := strconv.Atoi(obj.MakeYear)
	r.Year = int32(year)

	totalWeight, _ := strconv.Atoi(obj.MakeYear)
	r.TotalWeight = int32(totalWeight)

	ownWeight, _ := strconv.Atoi(obj.MakeYear)
	r.OwnWeight = int32(ownWeight)

	return r
}
