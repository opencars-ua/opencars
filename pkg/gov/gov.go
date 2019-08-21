package gov

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const DefaultHost = "https://data.gov.ua"

var BaseHost = DefaultHost

type Client struct {
	http *http.Client
}

func NewClient() *Client {
	api := new(Client)

	api.http = &http.Client{
		Timeout: 5 * time.Second,
	}

	return api
}

func (c *Client) get(path string) (*Response, error) {
	url := fmt.Sprintf("%s/%s", BaseHost, path)

	resp, err := c.http.Get(url)
	if err != nil {

	}

	JSON := new(Response)
	if err := json.NewDecoder(resp.Body).Decode(JSON); err != nil {
		return nil, err
	}

	return JSON, nil
}

func (c *Client) DataPackage(id string) (*Package, error) {
	path := fmt.Sprintf("/api/3/action/package_show?id=%s", id)
	res, err := c.get(path)
	if err != nil {
		return nil, err
	}

	pkg := new(Package)
	json.Unmarshal(res.Result, pkg)

	return pkg, nil
}

func (c *Client) DataResource(id string) (*Resource, error) {
	path := fmt.Sprintf("/api/3/action/resource_show?id=%s", id)
	res, err := c.get(path)
	if err != nil {
		return nil, err
	}

	resource := new(Resource)
	json.Unmarshal(res.Result, resource)

	return resource, nil
}
