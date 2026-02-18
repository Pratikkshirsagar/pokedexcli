package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMap(config *config) error {
	url := "https://pokeapi.co/api/v2/location-area"
	if config.Next != nil {
		url = *config.Next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	location := LocationAreasResp{}

	err = json.Unmarshal(body, &location)
	if err != nil {
		return err
	}

	config.Next = location.Next
	config.Previous = location.Previous

	for _, loc := range location.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}
