package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/json-iterator/go"

	"github.com/opencars/opencars/internal/database"
	"github.com/opencars/opencars/pkg/models"
	"github.com/opencars/opencars/pkg/translator"
)

var (
	// DB is an instance of Database Interface.
	DB   database.Adapter
	json = jsoniter.ConfigFastest
)

func Transport(w http.ResponseWriter, req *http.Request) {
	// start := time.Now()
	cars := make([]models.Transport, 0)
	number := translator.ToUA(req.FormValue("number"))
	limit := 1

	if strings.TrimSpace(number) == "" {
		http.Error(w, "number is empty", http.StatusBadRequest)
		return
	}

	if tmp, err := strconv.Atoi(req.FormValue("limit")); err == nil {
		limit = tmp
	}

	if err := DB.Select(&cars, limit, "number = ?", number); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	// fmt.Printf("Execution time: %s\n", time.Since(start))
}

// HealthHandler is a net/http handler for health checks.
func Health(w http.ResponseWriter, _ *http.Request) {
	if DB.Healthy() {
		msg := "database is not healthy"
		http.Error(w, msg, http.StatusServiceUnavailable)
	}
}

func Run() {
	http.HandleFunc("/transport", Transport)
	http.HandleFunc("/health", Health)

	fmt.Println("Listening port 8080")

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		panic(err.Error())
	}
}
