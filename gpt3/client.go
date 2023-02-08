package gpt3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Request is a struct to hold request data
type Request struct {
	Model     string `json:"model"`
	Prompt    string `json:"prompt"`
	MaxTokens int    `json:"max_tokens"`
}

// Response is a struct to hold response data
type Response struct {
	Choices []struct {
		Text string `json:"text"`
	} `json:"choices"`
}

func GetChoices(prompt string, apiKey string) (Response, error) {
	// Define endpoint URL
	url := "https://api.openai.com/v1/completions"

	// Create a new request
	request := Request{
		Model:     "text-davinci-003",
		Prompt:    prompt,
		MaxTokens: 256,
	}

	// Marshal the request struct into JSON
	requestJSON, err := json.Marshal(request)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	// Create a new POST request with the request JSON as the request body
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestJSON))
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// Create a new HTTP client and send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}
	defer res.Body.Close()

	// Read the response body and check for any errors
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	// Unmarshal the response JSON into a response struct
	var response Response
	err = json.Unmarshal(resBody, &response)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	return response, nil
}
