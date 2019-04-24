package main

import (
	"github.com/go-pg/pg"

	"github.com/opencars/opencars/internal/database"
	"github.com/opencars/opencars/internal/http"
)

// Adapter implements interface Adapter from database	 package.
type Adapter struct {
	db *pg.DB
}

func NewAdapter(db *pg.DB) *Adapter {
	adapter := new(Adapter)

	adapter.db = db

	return adapter
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

// Healthy performs application health check.
func (adapter *Adapter) Healthy() bool {
	_, err := adapter.db.Exec("SELECT 1")
	return err != nil
}

func main() {
	db := database.Must(database.DB())

	http.Storage = NewAdapter(db)
	http.Run(":8080")
}
