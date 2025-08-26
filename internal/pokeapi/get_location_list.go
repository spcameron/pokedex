package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type LocationAreaBatch struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationAreaBatch(inputURL *string) (LocationAreaBatch, error) {
	url := baseURL + "/location-area"
	if inputURL != nil {
		url = *inputURL
	}

	cachedData, exists := c.cache.Get(url)
	if exists {
		result := LocationAreaBatch{}
		err := json.Unmarshal(cachedData, &result)
		if err != nil {
			return LocationAreaBatch{}, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaBatch{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaBatch{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaBatch{}, err
	}

	result := LocationAreaBatch{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationAreaBatch{}, err
	}

	c.cache.Add(url, data)
	return result, nil

}
