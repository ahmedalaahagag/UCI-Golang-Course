package main

import (
	"fmt"
	"math"
	"strconv"
)

func main()  {
	fmt.Printf("Input acceleration :\n")
	a := takeInput()
	fmt.Printf("Input velocity :\n")
	v := takeInput()
	fmt.Printf("Input displacment :\n")
	s := takeInput()
	fn := GenDisplaceFn(a,v,s)
	fmt.Printf("Input time :\n")
	t := takeInput()
	fmt.Println(fn(t))
}

func takeInput() float64 {
	var inputString string
	var convertedNumber float64

	fmt.Scan(&inputString)

	convertedNumber, err := strconv.ParseFloat(inputString, 64)

	if err != nil {
		fmt.Printf("Input must be an intger \n")
	}
	return convertedNumber
}


func GenDisplaceFn (a, v , s float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return a*(math.Pow(t, 2)*0.5) + (v * t) + (s)
	}
	return fn
}