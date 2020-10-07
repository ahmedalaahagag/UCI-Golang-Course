package main

import (
	"fmt"
	"strconv"
	)

func main()  {
	var f float32 = 12.343
	fmt.Printf(strconv.Itoa(trunc(f)))
}


func trunc(f float32) int {
	return int(f)
}

