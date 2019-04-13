package main

import (
	"github.com/go-pg/pg"
	"log"

	"github.com/opencars-ua/opencars/internal/database"
	"github.com/opencars-ua/opencars/internal/http"
)

// Adapter implements interface Adapter from database	 package.
type Adapter struct {
	db *pg.DB
}

// Select returns set of objects searched by SQL SELECT.
func (adapter *Adapter) Select(
	model interface{},
	limit int,
	condition string,
	params ...interface{},
) error {
	query := adapter.db.Model(model).Where(condition, params...)
	return query.Limit(limit).Select()
}

// Select returns set of objects searched by SQL SELECT.
func (adapter *Adapter) Healthy() bool {
	_, err := adapter.db.Exec("SELECT 1")

	log.Printf("Database: %v\n", err)

	return err == nil
}
func main() {
	sql := database.Must(database.DB())

	http.DB = &Adapter{sql}
	http.Run()

	defer sql.Close()
}
