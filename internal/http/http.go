package http

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/json-iterator/go"

	"github.com/opencars/opencars/internal/database"
	"github.com/opencars/opencars/pkg/model"
	"github.com/opencars/opencars/pkg/translator"
)

// Error is error JSON format with error description.
type Error struct {
	Error string `json:"error"`
}

var (
	// Storage is an instance of Database Interface.
	Storage database.Adapter
	json    = jsoniter.ConfigFastest
)

func SendError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(Error{msg}); err != nil {
		log.Panic(err.Error())
	}
}

func Transport(w http.ResponseWriter, req *http.Request) {
	cars := make([]model.Operation, 0)
	number := translator.ToUA(req.FormValue("number"))
	limit := 1

	w.Header().Set("Server", "opencars")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if number == "" {
		SendError(w,http.StatusBadRequest, "number is empty")
		return
	}

	if tmp, err := strconv.Atoi(req.FormValue("limit")); err == nil {
		limit = tmp
	} else if req.FormValue("limit") != "" {
		SendError(w,http.StatusBadRequest, "limit is not valid")
		return
	}

	if err := Storage.Select(&cars, limit, "number = ?", number); err != nil {
		SendError(w,http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cars); err != nil {
		SendError(w,http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		log.Println(err)
		return
	}
}

// HealthHandler is a net/http handler for health checks.
func Health(w http.ResponseWriter, _ *http.Request) {
	if Storage.Healthy() {
		msg := "database is not healthy"
		http.Error(w, msg, http.StatusServiceUnavailable)
	}
}

func Run(addr string) {
	log.Printf("Server is listening %s\n", addr)

	router := http.NewServeMux()
	router.HandleFunc("/transport", Transport)
	router.HandleFunc("/health", Health)

	server := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	server.SetKeepAlivesEnabled(true)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not listen on %s. Error: %v\n", addr, err)
	}
}
