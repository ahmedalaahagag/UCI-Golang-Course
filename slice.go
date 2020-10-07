package main

import (
	"fmt"
	"sort"
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
		sort.Ints(sli)
		fmt.Printf("%v \n", sli)
	}
}
