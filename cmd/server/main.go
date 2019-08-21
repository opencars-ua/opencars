package main

import (
	"flag"
	"os"

	"github.com/go-pg/pg"

	"github.com/opencars/opencars/internal/config"
	"github.com/opencars/opencars/internal/http"
	"github.com/opencars/opencars/internal/storage"
)

// Adapter implements interface Adapter from storage package.
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
	return query.Order("id ASC").Limit(limit).Select()
}

// Select returns set of objects searched by SQL SELECT.
func (adapter *Adapter) Update(
	model interface{},
) error {
	return adapter.db.Update(model)
}

// Select returns set of objects searched by SQL SELECT.
func (adapter *Adapter) Insert(
	model interface{},
) error {
	_, err := adapter.db.Model(model).Insert()
	return err
}

// Healthy performs application health check.
func (adapter *Adapter) Healthy() bool {
	_, err := adapter.db.Exec("SELECT 1")
	return err != nil
}

var (
	path = flag.String("config", "./config/opencars.toml", "Path to configuration file")
)

func main() {
	flag.Parse()

	// Load configuration.
	config, err := config.New(*path)
	if err != nil {
		panic(err)
	}

	// Create database connection.
	db, err := storage.New(config)
	if err != nil {
		panic(err)
	}

	// Initialise database connection.
	err = storage.Migrate(db)
	if err != nil {
		panic(err)
	}

	// Run web server.
	http.Storage = NewAdapter(db)
	http.Run(config.API.Address(), os.Getenv("REGS_BASE_URL"))
}
