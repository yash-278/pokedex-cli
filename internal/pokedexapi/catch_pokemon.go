package pokedexapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (catchStatus Pokemon, err error) {
	url := baseURL + "/pokemon/" + pokemonName

	pokemonResp := Pokemon{}

	if val, ok := c.cache.Get(pokemonName); ok {
		err = json.Unmarshal(val, &pokemonResp)
		if err != nil {
			return pokemonResp, err
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return pokemonResp, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return pokemonResp, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return pokemonResp, err
	}

	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return pokemonResp, err
	}

	return pokemonResp, nil
}
