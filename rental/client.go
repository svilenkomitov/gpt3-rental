package rental

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Rental is a struct to hold rental data
type Rental struct {
	Data struct {
		Id         string `json:"id"`
		Attributes struct {
			CancelPolicy string `json:"cancel_policy"`
			Catalog      struct {
				ProcessedAmenities struct {
					Bathroom []struct {
						Available bool   `json:"available"`
						Name      string `json:"name"`
					} `json:"bathroom"`
					Electric []struct {
						Available bool   `json:"available"`
						Name      string `json:"name"`
					} `json:"electric"`
					Entertainment []struct {
						Available bool   `json:"available"`
						Name      string `json:"name"`
					} `json:"entertainment"`
					Kitchen []struct {
						Available bool   `json:"available"`
						Name      string `json:"name"`
					} `json:"kitchen"`
					Other []struct {
						Available bool   `json:"available"`
						Name      string `json:"name"`
					} `json:"other"`
					Temperature []struct {
						Available bool   `json:"available"`
						Name      string `json:"name"`
					} `json:"temperature"`
				} `json:"processed_amenities"`
			} `json:"catalog"`
			ChildrenCount      int    `json:"children_count"`
			DisplayVehicleType string `json:"display_vehicle_type"`
			Location           struct {
				City    string  `json:"city"`
				Country string  `json:"country"`
				County  string  `json:"county"`
				Lat     float64 `json:"lat"`
				Lng     float64 `json:"lng"`
				State   string  `json:"state"`
				Zip     string  `json:"zip"`
			} `json:"location"`
			Name         string `json:"name"`
			PricePerDay  int    `json:"price_per_day"`
			Sleeps       int    `json:"sleeps"`
			SleepsAdults int    `json:"sleeps_adults"`
			SleepsKids   int    `json:"sleeps_kids"`
			VehicleMake  string `json:"vehicle_make"`
			VehicleModel string `json:"vehicle_model"`
			VehicleTitle string `json:"vehicle_title"`
		} `json:"attributes"`
	} `json:"data"`
}

func GetRental(id int) (Rental, error) {
	// Define the API endpoint URL
	url := fmt.Sprintf("https://search.staging.outdoorsy.com/rentals/%d", id)

	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return Rental{}, nil
	}

	// Create a new HTTP client and send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Rental{}, nil
	}
	defer res.Body.Close()

	// Read the response body and check for any errors
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Rental{}, nil
	}

	// Unmarshal the response JSON into a rental struct
	var rental Rental
	err = json.Unmarshal(resBody, &rental)
	if err != nil {
		fmt.Println(err)
		return Rental{}, nil
	}

	return rental, nil
}
