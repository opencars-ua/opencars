package server

import (
	"fmt"
	"net/http"

	"github.com/shal/opencars/internal/database"
	"github.com/shal/opencars/pkg/handlers"
)

func Run() {
	DB := database.Must(database.DB())
	mux := http.NewServeMux()

	mux.Handle("/transport", handlers.TransportSearchHandler{DB: DB})

	fmt.Println("Listening port 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
