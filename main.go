package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sli []int

	for {
		var inputString string
		var convertedNumber int

		fmt.Scan(&inputString)

		if inputString[0] == 'X' || inputString[0] == 'x' {
			return
		}

		convertedNumber, err := strconv.Atoi(inputString)

		if err != nil {
			fmt.Printf("Input must be an intger \n")
			continue
		}

		sli = append(sli, convertedNumber)
		fmt.Printf("%v \n", sli)
	}
}
