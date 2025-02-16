package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (Cow) Eat() string {
	return "grass"
}

func (Cow) Move() string {
	return "walk"
}

func (Cow) Speak() string {
	return "moo"
}

func (Bird) Eat() string {
	return "worms"
}

func (Bird) Move() string {
	return "fly"
}

func (Bird) Speak() string {
	return "peep"
}

func (Snake) Eat() string {
	return "mice"
}

func (Snake) Move() string {
	return "slither"
}

func (Snake) Speak() string {
	return "hsss"
}

func getAnimalByTypeName(typeName string) Animal {
	switch strings.ToLower(typeName) {
	case "cow":
		return Cow{}
	case "bird":
		return Bird{}
	case "snake":
		return Snake{}
	default:
		return nil
	}
}

func getAnimalAction(animal Animal, action string) func() string {
	switch strings.ToLower(action) {
	case "eat":
		return animal.Eat
	case "move":
		return animal.Move
	case "speak":
		return animal.Speak
	default:
		return nil
	}
}

func main() {
	animals := make(map[string]Animal)
	for {
		fmt.Print("query > ")
		var cmd, arg1, arg2 string
		_, err := fmt.Scan(&cmd, &arg1, &arg2)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}

		if cmd == "newanimal" {
			animal := getAnimalByTypeName(arg2)
			if animal == nil {
				fmt.Println("Unknown animal type")
				continue
			}
			animals[arg1] = animal
			fmt.Println("Created it!")
		} else if cmd == "query" {
			animal, ok := animals[arg1]
			if !ok {
				fmt.Println("Unknown animal")
				continue
			}
			animalAction := getAnimalAction(animal, arg2)
			if animalAction == nil {
				fmt.Println("Unknown action")
				continue
			}
			fmt.Println(animalAction())
		} else {
			fmt.Println("Unknown command")
		}
	}
}
