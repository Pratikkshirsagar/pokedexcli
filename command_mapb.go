package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func commandMapB(config *config) error {
	if config.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *config.Previous
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
