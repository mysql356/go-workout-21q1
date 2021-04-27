package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Employee struct {
		Name    string //`json:"a"`
		Address string `json:"b,omitempty"`
		Draft   bool
	}
	e := Employee{"aa", "", true}

	data, err := json.Marshal(e)
	//data, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	////////////

}
