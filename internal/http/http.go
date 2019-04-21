package http

import (
	"github.com/json-iterator/go"
	"github.com/opencars/opencars/internal/database"
	"github.com/opencars/opencars/pkg/model"
	"github.com/opencars/opencars/pkg/translator"
	"log"
	"net/http"
	"strconv"
)

var (
	// Storage is an instance of Database Interface.
	Storage database.Adapter
	json    = jsoniter.ConfigFastest
)

func Transport(w http.ResponseWriter, req *http.Request) {
	cars := make([]model.Transport, 0)
	number := translator.ToUA(req.FormValue("number"))
	limit := 1

	if number == "" {
		http.Error(w, "number is empty", http.StatusBadRequest)
		return
	}

	if tmp, err := strconv.Atoi(req.FormValue("limit")); err == nil {
		limit = tmp
	}

	if err := Storage.Select(&cars, limit, "number = ?", number); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
}

// HealthHandler is a net/http handler for health checks.
func Health(w http.ResponseWriter, _ *http.Request) {
	if Storage.Healthy() {
		msg := "database is not healthy"
		http.Error(w, msg, http.StatusServiceUnavailable)
	}
}

func Run(addr string) {
	http.HandleFunc("/transport", Transport)
	http.HandleFunc("/health", Health)

	log.Printf("Server is listening %s\n", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err.Error())
	}
}
