package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Enter a String: ")
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	var low string = strings.ToLower(input)

	//if conditions
	var prefix bool = strings.HasPrefix(low, "i")
	var suffix bool = strings.HasSuffix(low, "n\n")
	//var suffix bool = low[length-2:length-1] == "n"
	var hasa bool = strings.Contains(low, "a")
	if hasa && prefix && suffix {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
