package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/go-pg/pg"

	"github.com/opencars/opencars/internal/config"
	"github.com/opencars/opencars/internal/http"
	"github.com/opencars/opencars/internal/storage"
)

// Adapter implements interface Adapter from database package.
type Adapter struct {
	db *pg.DB
}

//
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

//
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
	path    = flag.String("config", "./config/opencars.toml", "Path to configuration file")
	prefix  = flag.String("prefix", "", "First letters of unique document ID")
	seed    = flag.Int64("seed", time.Now().UTC().UnixNano(), "Pseudo-random number")
	threads = flag.Uint("threads", 100, "Number of threads")
)

func main() {
	flag.Parse()

	rand.Seed(*seed)

	// Load configuration.
	conf, err := config.New(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Create database connection.
	db, err := storage.New(conf)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	index := sort.SearchStrings(conf.HSC.Prefixes, *prefix)
	if index > len(conf.HSC.Prefixes) || index < 0 {
		fmt.Fprintln(os.Stderr, "Prefix is not valid!")
		os.Exit(1)
	}

	http.Storage = NewAdapter(db)
	amount := *threads
	http.VINHacker(*prefix, conf.HSC.URL(), uint16(amount))
}
