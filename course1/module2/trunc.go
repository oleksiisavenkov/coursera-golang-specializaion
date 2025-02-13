/* Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
 */
package main

import (
	"fmt"
	"math"
)

func main() {
	var input float64
	fmt.Print("Enter a floating point number: ")
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(math.Trunc(input))
}
