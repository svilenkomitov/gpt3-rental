package gpt3_rental

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/svilenkomitov/gpt3-rental/gpt3"
	"github.com/svilenkomitov/gpt3-rental/nps"
	"github.com/svilenkomitov/gpt3-rental/rental"
	"net/http"
	"os"
	"strconv"
)

type Response struct {
	Description string `json:"description"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Start() {
	// Load the environment variables from .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Get the value of an environment variable
	gpt3ApiKey := os.Getenv("GPT3_API_KEY")
	fmt.Println("GPT3 API Key:", gpt3ApiKey)

	// Get the value of an environment variable
	npsApiKey := os.Getenv("NPS_API_KEY")
	fmt.Println("NPS API Key:", npsApiKey)

	r := mux.NewRouter()
	r.HandleFunc("/{id}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			errRes := ErrorResponse{Error: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errRes)
			return
		}

		rental, err := rental.GetRental(id)
		if err != nil {
			errRes := ErrorResponse{Error: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errRes)
			return
		}

		campgrounds, err := nps.GetCampgrounds(rental.Data.Attributes.Location.State, npsApiKey)
		if err != nil {
			errRes := ErrorResponse{Error: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errRes)
			return
		}

		prompt := fmt.Sprintf("create description for rv rental using the following data %v. You can also add information for near campgrounds %v", rental, campgrounds)
		resp, err := gpt3.GetChoices(prompt, gpt3ApiKey)
		if err != nil {
			errRes := ErrorResponse{Error: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(errRes)
			return
		}

		res := Response{Description: resp.Choices[0].Text}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(res)
	})

	err = http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println(err)
	}
}
