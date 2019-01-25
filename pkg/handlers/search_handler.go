package handlers

import (
	"fmt"
	"github.com/shal/opencars/pkg/translit"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-pg/pg"
	"github.com/json-iterator/go"
	"github.com/shal/opencars/pkg/models"
)

type TransportSearchHandler struct {
	DB *pg.DB
}

func (h TransportSearchHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	start := time.Now()

	cars := make([]models.Transport, 0)
	json := jsoniter.ConfigCompatibleWithStandardLibrary

	number := translit.ToUA(req.FormValue("number"))
	limit := req.FormValue("limit")

	if strings.TrimSpace(number) == "" {
		http.Error(w, "number is empty", http.StatusBadRequest)
		return
	}

	query := h.DB.Model(&cars).Where("number LIKE ?", number)
	if res, err := strconv.Atoi(limit); err != nil {
		query.Limit(res).Select()
	} else {
		query.Select()
	}

	w.Header().Set("Content-Type", "application/json")

	bytes, err := json.Marshal(cars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)

	fmt.Printf("Execution time: %s\n", time.Since(start))
}
