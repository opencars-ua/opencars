package http

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/json-iterator/go"
	"github.com/opencars-ua/opencars/pkg/models"
	"github.com/opencars-ua/opencars/pkg/translator"
)

var (
	DB   Database
	json = jsoniter.ConfigFastest
)

// Database interface makes handler testable.
type Database interface {
	Select(
		model interface{},
		limit int,
		condition string,
		params ...interface{},
	) error
}

func Handler(w http.ResponseWriter, req *http.Request) {
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

func Run() {
	http.HandleFunc("/transport", Handler)

	fmt.Println("Listening port 8080")

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		panic(err.Error())
	}
}
