package pokeapi

import (
	"encoding/json"
	"io"
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

	if val, ok := c.cache.Get(reqURL); ok {
		locationsResp := RespLocationAreas{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespLocationAreas{}, err
		}

		return locationsResp, nil
	}

	resp, err := http.Get(reqURL)
	if err != nil {
		return RespLocationAreas{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespLocationAreas{}, err
	}

	var response RespLocationAreas
	err = json.Unmarshal(data, &response)
	if err != nil {
		return RespLocationAreas{}, err
	}

	c.cache.Add(reqURL, data)
	return response, nil
}
