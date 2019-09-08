package config

import (
	"log"
	"sort"
	"strconv"

	"github.com/BurntSushi/toml"
)

// Settings is decoded configuration file.
type Settings struct {
	Database Database `toml:"database"`
	HSC      HSC      `toml:"hsc"`
	API      API      `toml:"api"`
}

// Database contains configuration details for database.
type Database struct {
	Network    string `toml:"network"`
	Host       string `toml:"host"`
	Port       int    `toml:"port"`
	User       string `toml:"username"`
	Password   string `toml:"password"`
	Name       string `toml:"database"`
	MaxRetries int    `toml:"max_retries"`
	Pool       int    `toml:"pool"`
}

// HSC contains configuration details for third-party website.
type HSC struct {
	Host     string   `toml:"host"`
	Enabled  bool     `toml:"enabled"`
	Prefixes []string `toml:"prefixes"`
}

// API helps configure API host and port.
type API struct {
	Host string `tonl:"host"`
	Port int    `toml:"port"`
}

// Address return API address in "host:port" format.
func (db *Database) Address() string {
	return db.Host + ":" + strconv.Itoa(db.Port)
}

// URL returns complete URL for third-party website.
func (hsc *HSC) URL() string {
	return hsc.Host
}

// Address returns API complete address.
func (api *API) Address() string {
	return api.Host + ":" + strconv.Itoa(api.Port)
}

// New reads application configuration from specified file path.
func New(path string) (*Settings, error) {
	config := new(Settings)
	if _, err := toml.DecodeFile(path, config); err != nil {
		log.Fatal(err)
	}

	// Sort array of strings.
	sort.Strings(config.HSC.Prefixes)

	return config, nil
}
