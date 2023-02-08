package main

import (
	"fmt"
	"github.com/svilenkomitov/gpt3-rental/gpt3"
)

func main() {
	resp, err := gpt3.GetChoices("Say hi", "******")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.Choices[0].Text)
}
