package apiserver

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	"github.com/gorilla/schema"
	jsoniter "github.com/json-iterator/go"

	"github.com/opencars/opencars/internal/storage"
	"github.com/opencars/opencars/pkg/version"
)

// Error is error JSON format with error description.
type Error struct {
	Error string `json:"error"`
}

var (
	// Storage is an instance of storage.Base{} Interface.
	Storage storage.Base
	json    = jsoniter.ConfigFastest
)

var (
	// ErrInvalidNumber is an error for notifying about number invalidity.
	ErrInvalidNumber = errors.New("invalid number")
	// ErrInvalidCode is an error for notifying about code invalidity.
	ErrInvalidCode = errors.New("invalid code")
	// ErrRemoteBroken is an error for notifying about remote problems.
	ErrRemoteBroken = errors.New("remote server is not available")
	// ErrInternal is an error for notifying about internal problems.
	ErrInternal = errors.New(http.StatusText(http.StatusInternalServerError))
	// ErrNotFound is an error for notifying that entity does not exist.
	ErrNotFound = errors.New("not found")
)

var decoder = schema.NewDecoder()

func sendError(w http.ResponseWriter, code int, msg string) {
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(Error{msg}); err != nil {
		log.Panic(err.Error())
	}
}

// Server is a main server middleware.
// Adds application headers.
func Server(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Server", "opencars")
		w.Header().Set("Connection", "keep-alive")
		w.Header().Set("Content-Type", "application/json; charset=utf-8")

		handler.ServeHTTP(w, req)
	}
}

// Validator validates request.
type Validator interface {
	Validate(r *http.Request) error
}

func decodeAndValidate(r *http.Request, v Validator) error {
	if err := decoder.Decode(v, r.URL.Query()); err != nil {
		return err
	}

	return v.Validate(r)
}

// HealthHandler is a net/http handler for health checks.
func health(w http.ResponseWriter, _ *http.Request) {
	if !Storage.Healthy() {
		msg := "database is not healthy"
		http.Error(w, msg, http.StatusServiceUnavailable)
	}
}

// Run registers routes and starts web server.
func Run(addr, uri string) {
	log.Printf("Server is listening %s\n", addr)

	router := mux.NewRouter()
	vehicle := router.PathPrefix("/vehicle/").Methods("GET", "OPTIONS").Subrouter()
	// GET /vehicle/registrations.
	vehicle.Handle("/registrations", newRegsHandler(uri))
	// GET /vehicle/operations.
	vehicle.HandleFunc("/operations", operations)

	// GET /health.
	router.HandleFunc("/health", health)
	// GET /version.
	router.Handle("/version", version.Handler{})

	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"Api-Key"})

	cors := handlers.CORS(origins, methods, headers)(Server(router))
	server := &http.Server{
		Addr:         addr,
		Handler:      cors,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	server.SetKeepAlivesEnabled(true)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not listen on %s. Error: %v\n", addr, err)
	}
}
