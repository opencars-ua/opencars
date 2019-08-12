package hsc

import (
	"fmt"
	"net/http"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

type Registration struct {
	Brand        string `json:"brand"`
	Capacity     string `json:"capacity"`
	Color        string `json:"color"`
	DFirstReg    string `json:"dFirstReg"`
	DReg         string `json:"dReg"`
	Fuel         string `json:"fuel"`
	Kind         string `json:"kind"`
	MakeYear     string `json:"makeYear"`
	Model        string `json:"model"`
	NDoc         string `json:"nDoc"`
	NRegNew      string `json:"nRegNew"`
	OwnWeight    string `json:"ownWeight"`
	RankCategory string `json:"rankCategory"`
	SDoc         string `json:"sDoc"`
	TotalWeight  string `json:"totalWeight"`
	VIN          string `json:"vin"`
}

type API struct {
	baseUrl string
}

func New(uri string) *API {
	api := new(API)

	api.baseUrl = uri

	return api
}

func (api *API) VehiclePassport(code string) ([]Registration, error) {
	uri := fmt.Sprintf(
		"%s/gateway-edrmvs/api/verification/spr/%s/%s",
		api.baseUrl, code[:3], code[3:],
	)

	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	info := make([]Registration, 0)
	if err := json.NewDecoder(resp.Body).Decode(&info); err != nil {
		return nil, err
	}

	return info, nil
}
