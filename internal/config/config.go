package config

import (
	"log"
	"sort"
	"strconv"

	"github.com/BurntSushi/toml"
)

//
type TOML struct {
	Database Database `toml:"database"`
	HSC      HSC      `toml:"hsc"`
	API      API      `toml:"api"`
}

//
type Database struct {
	Network     string `toml:"network"`
	Host        string `toml:"host"`
	Port        int    `toml:"port"`
	User        string `toml:"username"`
	Password    string `toml:"password"`
	Name        string `toml:"database"`
	Max_retries int    `toml:"max_retries"`
	Pool        int    `toml:"pool"`
}

//
type HSC struct {
	Host     string   `toml:"host"`
	Enabled  bool     `toml:"enabled"`
	Prefixes []string `toml:"prefixes"`
}

//
type API struct {
	Host string `tonl:"host"`
	Port int    `toml:"port"`
}

// Address return API address in "host:port" format.
func (db *Database) Address() string {
	return db.Host + ":" + strconv.Itoa(db.Port)
}

//
func (hsc *HSC) URL() string {
	return hsc.Host
}

// Address return API address in "host:port" format.
func (api *API) Address() string {
	return api.Host + ":" + strconv.Itoa(api.Port)
}

// New reads application configuration from specified file path.
func New(path string) (*TOML, error) {
	config := new(TOML)
	if _, err := toml.DecodeFile(path, config); err != nil {
		log.Fatal(err)
	}

	// Sort array of strings.
	sort.Strings(config.HSC.Prefixes)

	return config, nil
}
