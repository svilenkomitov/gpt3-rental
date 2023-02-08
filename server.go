package gpt3_rental

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/svilenkomitov/gpt3-rental/gpt3"
	"github.com/svilenkomitov/gpt3-rental/rental"
	"net/http"
	"strconv"
)

type Response struct {
	Description string `json:"description"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func Start() {
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

		prompt := fmt.Sprintf("create description for rv rental using the following data %v", rental)
		resp, err := gpt3.GetChoices(prompt, "*******")
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

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println(err)
	}
}
