package main

import (
	"fmt"
	"math"
)

func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + s
	}
}

func main() {
	var a, v, s, t float64
	fmt.Print("Enter a, v, s, t: ")
	_, err := fmt.Scan(&a, &v, &s, &t)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fn := GenDisplaceFn(a, v, s)
	fmt.Println(fn(t))

}
