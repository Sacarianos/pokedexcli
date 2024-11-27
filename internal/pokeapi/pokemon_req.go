package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	endpoint := "/location-area"
	fullURL := baseURL + endpoint
	if pageURL != nil {
		fullURL = *pageURL
	}

	//check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		LocationAreasRes := LocationAreasResponse{}
		err := json.Unmarshal(data, &LocationAreasRes)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return LocationAreasRes, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	LocationAreasRes := LocationAreasResponse{}
	err = json.Unmarshal(data, &LocationAreasRes)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return LocationAreasRes, nil

}

func (c *Client) GetLocationArea(LocationAreaName string) (LocationArea, error) {
	endpoint := "/location-area/" + LocationAreaName
	fullURL := baseURL + endpoint

	//check the cache
	data, ok := c.cache.Get(fullURL)
	if ok {
		// cache hit
		LocationAreaRes := LocationArea{}
		err := json.Unmarshal(data, &LocationAreaRes)
		if err != nil {
			return LocationArea{}, err
		}
		return LocationAreaRes, nil
	}

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return LocationArea{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer res.Body.Close()

	data, err = io.ReadAll(res.Body)
	if err != nil {
		return LocationArea{}, err
	}

	LocationAreaRes := LocationArea{}
	err = json.Unmarshal(data, &LocationAreaRes)
	if err != nil {
		return LocationArea{}, err
	}

	c.cache.Add(fullURL, data)

	return LocationAreaRes, nil

}
