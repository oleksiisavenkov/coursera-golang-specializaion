package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Println("Enter a name and an address")
	_, err := fmt.Scan(&name, &address)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	person := map[string]string{"name": name, "address": address}
	jsonPerson, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(string(jsonPerson))

}
