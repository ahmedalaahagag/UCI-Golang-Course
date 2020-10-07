package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animals struct{ Food, Locomotion, Noise string }

func (a Animals) Eat() string {
	return a.Food
}
func (a Animals) Move() string {
	return a.Locomotion
}
func (a Animals) Speak() string {
	return a.Noise
}

func main()  {
	var animalsSlice = make(map[string]Animals)

	for {
		fmt.Printf("Select command (newanimal, query) :\n")
		commandParts := strings.Fields(TakeUserInput())

		_, found := FindInSlice([]string{"newanimal", "query"}, strings.ToLower(commandParts[0]))
		if !found {
			fmt.Println("Command not found")
			continue
		}

		if len(commandParts) < 3 {
			fmt.Println("Command args not correct")
			continue
		}

		switch commandParts[0] {
		case "newanimal":
			NewAnimalCommand(animalsSlice, commandParts)
			break
		case "query":
			QueryCommand(animalsSlice, commandParts)
			break
		}
	}
}

func NewAnimalCommand(animalSlice map[string]Animals, commandParts []string)  {
	switch strings.ToLower(commandParts[2]) {
	case "cow":
		animalSlice[commandParts[1]] = Animals{Noise: "mooo", Locomotion: "walk", Food: "grass"}
		break
	case "bird":
		animalSlice[commandParts[1]] = Animals{Noise: "peep", Locomotion: "fly", Food: "worms"}
		break
	case "snake":
		animalSlice[commandParts[1]] = Animals{Noise: "hsss ", Locomotion: "slither", Food: "mice"}
		break
	default:
		fmt.Println("Animal type not found")
		return
	}

	fmt.Println("Created it !")
}

func QueryCommand(animalSlice map[string]Animals, commandParts []string)  {
	if animal, found := animalSlice[commandParts[1]]; found {
		switch strings.ToLower(commandParts[2]) {
		case "eat":
			fmt.Println(animal.Eat())
			break
		case "move":
			fmt.Println(animal.Move())
			break
		case "speak":
			fmt.Println(animal.Speak())
			break
		default:
			fmt.Println("Method not found")
			return
		}
	} else {
		fmt.Println("No Animal with this name")
	}
}

func TakeUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	command, _ := reader.ReadString('\n')
	return command
}

func FindInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}