package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(area string) (RespLocation, error) {
	url := baseURL + "/location-area/" + area

	// Check the cache
	dat, ok := c.cache.Get(url)
	if ok {
		resp := RespLocation{}
		err := json.Unmarshal(dat, &resp)
		return resp, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespLocation{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocation{}, err
	}

	c.cache.Add(url, body)

	locationResp := RespLocation{}
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return RespLocation{}, err
	}

	return locationResp, nil
}
