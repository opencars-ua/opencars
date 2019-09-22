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
	// BuildDate holds the build date of opencars.
	// BuildDate = "I don't remember exactly"
)

var json = jsoniter.ConfigFastest

// Handler expose version routes.
type Handler struct{}

// ServeHTTP serves HTTP.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	v := struct {
		Version string `json:"version"`
		// BuildDate string `json:"build_date"`
		Go string `json:"go"`
	}{
		Version: Version,
		// BuildDate: BuildDate,
		Go: runtime.Version(),
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("version: %v", err)
	}
}
