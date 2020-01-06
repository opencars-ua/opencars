package version

import (
	"log"
	"net/http"
	"runtime"

	jsoniter "github.com/json-iterator/go"
)

var (
	// Version holds the current version of opencars.
	Version = "dev"
)

var json = jsoniter.ConfigFastest

// Handler expose version routes.
type Handler struct{}

// ServeHTTP serves HTTP.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Version string `json:"version"`
		Go      string `json:"go"`
	}{
		Version: Version,
		Go:      runtime.Version(),
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("version: %v", err)
	}
}
