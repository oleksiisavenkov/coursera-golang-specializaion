package main

import (
	"fmt"
	"slices"
)

func appendToSortedSlice(sli []int, num int) []int {
	for i, v := range sli {
		if num < v {
			return append(sli[:i], append([]int{num}, sli[i:]...)...)
		}
	}
	return append(sli, num)
}

func main() {

	stopwords := []string{"X", "x"}

	sli := make([]int, 0, 3)

	for {
		var input string
		var num int
		fmt.Print("Enter an integer or X: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		if slices.Contains(stopwords, input) {
			break
		}
		_, err = fmt.Sscan(input, &num)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		sli = appendToSortedSlice(sli, num)
		fmt.Println(sli)
	}

}
