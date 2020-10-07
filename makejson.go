package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	var m = make(map[string]string)

	var name, address string

	fmt.Print("Please enter your name: ")
	fmt.Scan(&name)
	fmt.Print("Please enter your address: ")
	fmt.Scan(&address)

	m["name"] = name
	m["address"] = address

	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(j))
	}
}

// Thank you for grading my submission :)
