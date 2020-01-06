package main

import (
	"flag"
	"log"
	"sort"

	"github.com/opencars/opencars/internal/config"
	"github.com/opencars/opencars/internal/http"
	"github.com/opencars/opencars/internal/storage"
	"github.com/opencars/opencars/pkg/adapter"
)

var (
	path    = flag.String("config", "./config/opencars.toml", "Path to configuration file")
	prefix  = flag.String("prefix", "", "First letters of unique document ID")
	threads = flag.Uint("threads", 100, "Number of threads")
)

func main() {
	flag.Parse()

	// Load configuration.
	settings, err := config.New(*path)
	if err != nil {
		log.Fatal(err)
	}

	// Create database connection.
	db, err := storage.New(settings)
	if err != nil {
		log.Fatal(err)
	}

	index := sort.SearchStrings(settings.HSC.Prefixes, *prefix)
	if index > len(settings.HSC.Prefixes) || index < 0 {
		log.Fatal("Prefix is not valid!")
	}

	http.Storage = adapter.New(db)
	http.VINHacker(*prefix, settings.HSC.URL(), uint16(*threads))
}
