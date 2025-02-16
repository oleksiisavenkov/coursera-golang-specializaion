package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Swap(sli []int, i int) {
	sli[i], sli[i+1] = sli[i+1], sli[i]
}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli); i++ {
		for j := 0; j < len(sli)-1; j++ {
			if sli[j] > sli[j+1] {
				Swap(sli, j)
			}
		}
	}
}

func main() {
	fmt.Print("Enter up to 10 integers (space separated): ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	strings := strings.Split(line, " ")
	ints := make([]int, 0, 10)
	for _, s := range strings {
		var i int
		_, err = fmt.Sscan(s, &i)
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		ints = append(ints, i)
	}
	BubbleSort(ints)
	fmt.Println(ints)
}
