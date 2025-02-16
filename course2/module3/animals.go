package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	eat, move, speak string
}

var cow = Animal{eat: "grass", move: "walk", speak: "moo"}
var bird = Animal{eat: "worms", move: "fly", speak: "peep"}
var snake = Animal{eat: "mice", move: "slither", speak: "hsss"}

func (a Animal) Eat() string {
	return a.eat
}

func (a Animal) Move() string {
	return a.move
}

func (a Animal) Speak() string {
	return a.speak
}

func getAnimalByName(name string) Animal {
	switch strings.ToLower(name) {
	case "cow":
		return cow
	case "bird":
		return bird
	case "snake":
		return snake
	default:
		return Animal{}
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
	for {
		fmt.Print("Animal and Action > ")
		var name, action string
		_, err := fmt.Scan(&name, &action)
		if err != nil {
			fmt.Println("Error: ", err)
			continue
		}
		animal := getAnimalByName(name)
		animalAction := getAnimalAction(animal, action)
		if animalAction == nil {
			fmt.Println("Unknown action")
			continue
		}
		fmt.Println(animalAction())
	}
}
