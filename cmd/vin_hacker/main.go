package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"

	"github.com/opencars/opencars/pkg/adapter"

	"github.com/opencars/opencars/internal/config"
	"github.com/opencars/opencars/internal/http"
	"github.com/opencars/opencars/internal/storage"
)

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
	settings, err := config.New(*path)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	// Create database connection.
	db, err := storage.New(settings)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	index := sort.SearchStrings(settings.HSC.Prefixes, *prefix)
	if index > len(settings.HSC.Prefixes) || index < 0 {
		fmt.Fprintln(os.Stderr, "Prefix is not valid!")
		os.Exit(1)
	}

	http.Storage = adapter.New(db)
	amount := *threads
	http.VINHacker(*prefix, settings.HSC.URL(), uint16(amount))
}
