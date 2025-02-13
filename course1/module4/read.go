package main

import (
	"bufio"
	"fmt"
	"os"
)

type person struct {
	firstname, lastname string
}

func main() {
	var filename string
	people := make([]person, 0, 3)

	fmt.Print("Enter a filename: ")
	_, err := fmt.Scan(&filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		var p person
		fmt.Sscan(line, &p.firstname, &p.lastname)
		people = append(people, p)
	}

	for _, p := range people {
		fmt.Println(p)
	}

}
