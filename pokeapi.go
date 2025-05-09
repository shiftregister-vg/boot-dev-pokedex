package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getApiResponse(url string, v any) error {
	bytes, found := cache.Get(url)

	if !found {
		// make the http request to get lands
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		bytes, err = io.ReadAll(resp.Body)

		// cache the response
		cache.Add(url, bytes)
	}

	if err := json.Unmarshal(bytes, v); err != nil {
		fmt.Printf("could not unmarshal cached response [%s]: %v", url, err)
		return err
	}

	return nil
}

func getLocationAreas(limit, offset int) []*LocationArea {
	var respBody LocationAreasResp

	url := fmt.Sprintf("%s/location-area?limit=%d&offset=%d", baseURL, limit, offset)
	if err := getApiResponse(url, &respBody); err != nil {
		fmt.Printf("Could not get location areas from API: %v\n", err)
		return nil
	}

	return respBody.Results
}

func getLocationArea(name string) *LocationArea {
	var locationArea LocationArea

	url := fmt.Sprintf("%s/location-area/%s/", baseURL, name)
	if err := getApiResponse(url, &locationArea); err != nil {
		fmt.Printf("Could not get location area with name [%s]: %v\n", name, err)
		return nil
	}
	return &locationArea
}

func getPokemon(name string) *Pokemon {
	var pokemon Pokemon

	url := fmt.Sprintf("%s/pokemon/%s/", baseURL, name)
	if err := getApiResponse(url, &pokemon); err != nil {
		fmt.Printf("Could not get pokemon with name [%s]: %v", name, err)
		return nil
	}

	return &pokemon
}
