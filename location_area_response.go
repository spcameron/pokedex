package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type locationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) getLocationAreaBatch(inputURL *string) (locationAreaResponse, error) {
	url := "https://pokeapi.co/api/v2/location-area"
	if inputURL != nil {
		url = *inputURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return locationAreaResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResponse{}, err
	}

	result := locationAreaResponse{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return locationAreaResponse{}, err
	}

	return result, nil

}
