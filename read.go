package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	fmt.Print("Enter file name: \n")

	inputReader := bufio.NewReader(os.Stdin)
	fileName, _ := inputReader.ReadString('\n')

	file,err := os.Open(strings.TrimSuffix(fileName, "\n"))

	if err != nil {
		fmt.Print(err)
	}

	defer file.Close()

	type name struct {
		fname string
		lname string
	}
	var sli []name
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splitName := strings.Split(scanner.Text(), " ")
		tempName := name{
			fname: splitName[0],
			lname: splitName[1],
		}
		sli = append(sli,tempName)
	}

	for _,v := range sli {
		fmt.Printf("First name : %s , Last name : %s ",v.fname,v.lname)
	}
}