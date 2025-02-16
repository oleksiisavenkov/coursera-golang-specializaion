package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
	"sync"
)

func ReadArrayFromStdin() []int {
	fmt.Print("Enter integers (space separated): ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return nil
	}
	strings := strings.Split(line, " ")
	ints := make([]int, 0, 10)
	for _, s := range strings {
		var i int
		_, err = fmt.Sscan(s, &i)
		if err != nil {
			fmt.Println("Error: ", err)
			return nil
		}
		ints = append(ints, i)
	}
	return ints
}

func MergeSortedArrays(array1, array2 []int) []int {
	if len(array1) == 0 {
		return array2
	}
	if len(array2) == 0 {
		return array1
	}

	result := make([]int, len(array1)+len(array2))
	i, j := 0, 0
	for k := 0; k < len(result); k++ {
		if i >= len(array1) {
			copy(result[k:], array2[j:])
			break
		}
		if j >= len(array2) {
			copy(result[k:], array1[i:])
			break
		}
		if array1[i] < array2[j] {
			result[k] = array1[i]
			i++
		} else {
			result[k] = array2[j]
			j++
		}
	}
	return result
}

func MergeSortedArrayPairsParallel(arrayPairs [][2][]int) [][]int {

	ch := make(chan []int, len(arrayPairs))
	for i := 0; i < len(arrayPairs); i++ {
		go func(i int) {
			ch <- MergeSortedArrays(arrayPairs[i][0], arrayPairs[i][1])
		}(i)
	}

	result := make([][]int, 0, len(arrayPairs))
	for i := 0; i < len(arrayPairs); i++ {
		result = append(result, <-ch)
	}

	return result
}

func ToPairs(array [][]int) [][2][]int {
	pairs := make([][2][]int, int(math.Ceil(float64(len(array))/2)))
	for i := 0; i < len(array); i += 2 {
		if i+1 == len(array) {
			pairs[i/2] = [2][]int{array[i], {}}
		} else {
			pairs[i/2] = [2][]int{array[i], array[i+1]}
		}
	}
	return pairs
}

func MergeSortedArraysParallel(array [][]int) []int {
	for len(array) > 1 {
		pairs := ToPairs(array)
		array = MergeSortedArrayPairsParallel(pairs)
	}
	return array[0]
}

func ParallelSort(ints []int, parallelism int) {

	if parallelism < 1 {
		parallelism = 1
	}

	if len(ints) < parallelism {
		parallelism = len(ints)
	}

	chunkSize := len(ints) / parallelism

	wg := sync.WaitGroup{}
	wg.Add(parallelism)

	sorted_slices := make([][]int, parallelism)

	for i := 0; i < parallelism; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == parallelism-1 {
			end = len(ints)
		}
		go func() {
			sort.Ints(ints[start:end])
			sorted_slices[i] = ints[start:end]
			wg.Done()
		}()
	}

	wg.Wait()

	sorted_result := MergeSortedArraysParallel(sorted_slices)

	copy(ints, sorted_result)
}

func main() {
	ints := ReadArrayFromStdin()
	parallelism := 4
	ParallelSort(ints, parallelism)
	fmt.Println(ints)
}
