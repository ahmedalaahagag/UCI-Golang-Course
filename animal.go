package main

import (
	"fmt"
	"strings"
)

type Animal struct{ Food, Locomotion, Noise string }

func (a Animal) Eat() string {
	return a.Food
}
func (a Animal) Move() string {
	return a.Locomotion
}
func (a Animal) Speak() string {
	return a.Noise
}


func main()  {
	for {
		fmt.Printf("Select Animal (cow,bird,snake) :\n")
		animal := TakeInput()

		animals := []string{"cow", "bird", "snake"}
		_, found := Find(animals, strings.ToLower(animal))
		if !found {
			fmt.Println("Animal not found")
			continue
		}

		var userAnimal Animal

	switch strings.ToLower(animal) {
		case "cow":
			userAnimal.Food = "grass"
			userAnimal.Locomotion = "walk"
			userAnimal.Noise = "moo"
			break
		case "bird":
			userAnimal.Food = "worms"
			userAnimal.Locomotion = "fly"
			userAnimal.Noise = "peep"
			break
		case "snake":
			userAnimal.Food = "mice"
			userAnimal.Locomotion = "slither"
			userAnimal.Noise = "hsss"
			break
		}

		fmt.Printf("Input Action (Eat,Speak,Move) :\n")
		method := TakeInput()

		methods := []string{"eat", "move", "speak"}
		_, found = Find(methods, strings.ToLower(method))
		if !found {
			fmt.Println("Method not found")
			continue
		}

		switch strings.ToLower(method) {
		case "eat":
			fmt.Println(userAnimal.Eat())
			break
		case "move":
			fmt.Println(userAnimal.Move())
			break
		case "speak":
			fmt.Println(userAnimal.Speak())
			break
		}
	}
}

func TakeInput() string {
	var inputString string
	fmt.Scan(&inputString)
	return inputString
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}