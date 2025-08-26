package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaBatch(inputURL *string) (LocationAreaBatchResponse, error) {
	url := baseURL + "/location-area"
	if inputURL != nil {
		url = *inputURL
	}

	cachedData, exists := c.cache.Get(url)
	if exists {
		result := LocationAreaBatchResponse{}
		err := json.Unmarshal(cachedData, &result)
		if err != nil {
			return LocationAreaBatchResponse{}, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaBatchResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaBatchResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaBatchResponse{}, err
	}

	result := LocationAreaBatchResponse{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationAreaBatchResponse{}, err
	}

	c.cache.Add(url, data)
	return result, nil

}
