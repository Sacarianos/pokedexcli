package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(PokemonName string) (Pokemon, error) {
	endpoint := "/pokemon/" + PokemonName
	fullURL := baseURL + endpoint

	//check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		PokemonRes := Pokemon{}
		err := json.Unmarshal(data, &PokemonRes)
		if err != nil {
			return Pokemon{}, err
		}
		return PokemonRes, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	PokemonRes := Pokemon{}
	err = json.Unmarshal(data, &PokemonRes)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return PokemonRes, nil

}
