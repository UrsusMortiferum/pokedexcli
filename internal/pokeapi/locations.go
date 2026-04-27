package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (ResponseShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResponseShallowLocations{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return ResponseShallowLocations{},
			fmt.Errorf("Response failed with status code :%d \n", res.StatusCode)
	}

	var resShallowLocations ResponseShallowLocations
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&resShallowLocations); err != nil {
		return ResponseShallowLocations{}, err
	}

	return resShallowLocations, nil
}
