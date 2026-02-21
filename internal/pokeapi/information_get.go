package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetInformation(pokemon string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		pokemonInfoRes := Pokemon{}
		err := json.Unmarshal(val, &pokemonInfoRes)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonInfoRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonInfoRes := Pokemon{}
	err = json.Unmarshal(dat, &pokemonInfoRes)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	return pokemonInfoRes, nil
}
