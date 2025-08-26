package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationArea(input string) (LocationAreaResponse, error) {
	url := baseURL + "/location-area/" + input

	cachedData, exists := c.cache.Get(url)
	if exists {
		result := LocationAreaResponse{}
		err := json.Unmarshal(cachedData, &result)
		if err != nil {
			return LocationAreaResponse{}, err
		}

		return result, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreaResponse{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	result := LocationAreaResponse{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return LocationAreaResponse{}, err
	}

	c.cache.Add(url, data)
	return result, nil

}
