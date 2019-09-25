package model

// WantedTransport.
type WantedTransport struct {
	ID            string `json:"id"`
	OVD           string `json:"ovd"`
	Brand         string `json:"brand"`
	Model         string `json:"model"`
	Kind          string `json:"kind"`
	Color         string `json:"color"`
	Number        string `json:"number"`
	BodyNumber    string `json:"body_number"`
	ChassisNumber string `json:"chassis_number"`
	EngineNumber  string `json:"engine_number"`
	TheftDate     string `json:"theft_date"`
	InsertDate    string `json:"insert_date"`
}
