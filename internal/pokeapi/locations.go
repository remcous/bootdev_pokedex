package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreas struct {
	Count   int     `json:"count"`
	Next    *string `json:"next"`
	Prev    *string `json:"previous"`
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int   `json:"chance"`
				ConditionValues []any `json:"condition_values"`
				MaxLevel        int   `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationList(pageURL *string) (LocationAreas, error) {
	reqUrl := baseURL + "location-area"
	if pageURL != nil {
		reqUrl = *pageURL
	}

	if val, ok := c.cache.Get(reqUrl); ok {
		locationsResp := LocationAreas{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return LocationAreas{}, err
		}

		return locationsResp, nil
	}

	resp, err := http.Get(reqUrl)
	if err != nil {
		return LocationAreas{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	var response LocationAreas
	err = json.Unmarshal(data, &response)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cache.Add(reqUrl, data)
	return response, nil
}

func (c *Client) GetLocation(locationName string) (Location, error) {
	if locationName == "" {
		return Location{}, fmt.Errorf("location name cannot be empty")
	}

	reqUrl := baseURL + "location-area/" + locationName
	if val, ok := c.cache.Get(reqUrl); ok {
		locationResp := Location{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return Location{}, err
		}

		return locationResp, nil
	}

	resp, err := http.Get(reqUrl)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	var response Location
	err = json.Unmarshal(data, &response)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(reqUrl, data)
	return response, nil
}
