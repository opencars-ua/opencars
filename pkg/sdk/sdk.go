package sdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/opencars/opencars/pkg/model"
)

type Transport = model.Transport

type Client struct {
	uri string
}

func New(uri string) *Client {
	client := new(Client)

	client.uri = uri

	return client
}

func (client *Client) Search(number string) ([]Transport, error) {
	if number == "" {
		return nil, errors.New("number is empty")
	}

	query := client.uri + "/transport/?number=" + number
	response, err := http.Get(query)

	if err != nil {
		return nil, err
	}

	models := make([]Transport, 0)
	if err := json.NewDecoder(response.Body).Decode(&models); err != nil {
		return nil, errors.New("invalid response body")
	}

	return models, nil
}
