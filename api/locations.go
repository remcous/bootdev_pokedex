package api

import (
	"encoding/json"
	"net/http"
)

type RespLocationAreas struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c *Client) GetLocationList(pageURL *string) (RespLocationAreas, error) {
	reqURL := baseURL + "location-area"
	if pageURL != nil {
		reqURL = *pageURL
	}

	resp, err := http.Get(reqURL)
	if err != nil {
		return RespLocationAreas{}, err
	}
	defer resp.Body.Close()

	var response RespLocationAreas
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return RespLocationAreas{}, err
	}

	return response, nil
}
