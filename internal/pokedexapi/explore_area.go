package pokedexapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// Explore Area -
func (c *Client) ExploreArea(locationName string) (RespFullLocations, error) {
	url := baseURL + "/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespFullLocations{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespFullLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespFullLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespFullLocations{}, err
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespFullLocations{}, err
	}

	fullLocationResp := RespFullLocations{}
	err = json.Unmarshal(data, &fullLocationResp)

	if err != nil {
		return RespFullLocations{}, nil
	}

	c.cache.Add(url, data)
	return fullLocationResp, nil
}
