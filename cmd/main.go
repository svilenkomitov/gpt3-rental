package main

import (
	"fmt"
	"github.com/svilenkomitov/gpt3-rental/gpt3"
	"github.com/svilenkomitov/gpt3-rental/rental"
)

func main() {
	rental, err := rental.GetRental(140356)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(rental)

	resp, err := gpt3.GetChoices("Say hi", "********")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Choices[0].Text)
}
