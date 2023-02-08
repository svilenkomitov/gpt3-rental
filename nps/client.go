package nps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Campgrounds struct {
	Data []struct {
		Name        string `json:"name"`
		ParkCode    string `json:"parkCode"`
		Description string `json:"description"`
	} `json:"data"`
}

func GetCampgrounds(state string, apiKey string) (Campgrounds, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Build the request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://developer.nps.gov/api/v1/campgrounds?api_key=%s&stateCode=%s", apiKey, state), nil)
	if err != nil {
		fmt.Println(err)
		return Campgrounds{}, err
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Campgrounds{}, err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return Campgrounds{}, err
	}

	// Unmarshal the response JSON into a rental struct
	var campgrounds Campgrounds
	err = json.Unmarshal(body, &campgrounds)
	if err != nil {
		fmt.Println(err)
		return Campgrounds{}, err
	}

	return campgrounds, nil
}
