package gov

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// DefaultHost is a default host of government public data registry.
const DefaultHost = "https://data.gov.ua"

// BaseHost is customizable base host of government public data registry.
var BaseHost = DefaultHost

// Client is a government public data registry client.
type Client struct {
	http *http.Client
}

// New create new instance of public data registry client.
func New() *Client {
	api := new(Client)

	api.http = &http.Client{
		Timeout: 5 * time.Second,
	}

	return api
}

func (c *Client) get(path string) (*Result, error) {
	url := fmt.Sprintf("%s/%s", BaseHost, path)

	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}

	JSON := new(Result)
	if err := json.NewDecoder(response.Body).Decode(JSON); err != nil {
		return nil, err
	}

	return JSON, nil
}

// DataPackage returns latest information about package with ID.
// Makes GET request to public data registry.
func (c *Client) DataPackage(id string) (*Package, error) {
	path := fmt.Sprintf("/api/3/action/package_show?id=%s", id)

	response, err := c.get(path)
	if err != nil {
		return nil, err
	}

	pkg := new(Package)
	if err := json.Unmarshal(response.Result, pkg); err != nil {
		return nil, err
	}

	return pkg, nil
}
