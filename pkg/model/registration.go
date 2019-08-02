package model

import (
	"strconv"
	"strings"

	"github.com/opencars/opencars/pkg/hsc"
)

// Registration represents information from vehicle registration document.
type Registration struct {
	ID          int    `json:"-"            sql:"id,pk"`
	Brand       string `json:"brand"        sql:"brand"`
	Capacity    int    `json:"capacity"     sql:"capacity"`
	Color       string `json:"color"        sql:"color"`
	FirstReg    string `json:"first_reg"    sql:"first_reg"`
	Date        string `json:"date"         sql:"date"`
	Fuel        string `json:"fuel"         sql:"fuel"`
	Kind        string `json:"kind"         sql:"kind"`
	Body        string `json:"body"         sql:"body"`
	Year        int    `json:"year"         sql:"year"`
	Model       string `json:"model"        sql:"model"`
	Code        string `json:"code"         sql:"code"`
	Number      string `json:"number"       sql:"number"`
	TotalWeight int    `json:"total_weight" sql:"total_weight"`
	OwnWeight   int    `json:"own_weight"   sql:"own_weight"`
	Category    string `json:"category"     sql:"category"`
	VIN         string `json:"vin"          sql:"vin"`
}

// NewOperation parses CSV line into operation model.
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
	r.VIN = obj.Vin

	r.Capacity, _ = strconv.Atoi(obj.Capacity)
	r.Year, _ = strconv.Atoi(obj.MakeYear)
	r.TotalWeight, _ = strconv.Atoi(obj.TotalWeight)
	r.OwnWeight, _ = strconv.Atoi(obj.OwnWeight)

	return r
}
