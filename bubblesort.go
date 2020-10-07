package main

import (
	"fmt"
	"strconv"
)

func main()  {
	var sli []int
	for len(sli) < 10 {
		var inputString string
		var convertedNumber int

		fmt.Scan(&inputString)

		convertedNumber, err := strconv.Atoi(inputString)

		if err != nil {
			fmt.Printf("Input must be an intger \n")
			continue
		}

		sli = append(sli, convertedNumber)
	}
	BubbleSort(sli)
	fmt.Printf("%v \n", sli)
}

func BubbleSort(toSort []int) []int {
	for i := len(toSort); i > 0; i-- {
		for j := 1; j < i; j++ {
			if toSort[j-1] > toSort[j] {
				Swap(j, toSort)
			}
		}
	}
	return toSort
}
func Swap(index int, toSort []int)  {
	intermediate := toSort[index]
	toSort[index] = toSort[index-1]
	toSort[index-1] = intermediate
}