package main

import (
	"sort"
	"testing"
)

func runParallelSortTest(ints []int, parallelism int, t *testing.T) {
	ParallelSort(ints, parallelism)
	if !sort.IntsAreSorted(ints) {
		t.Errorf("Expected sorted array, got %v; parallelism: %v", ints, parallelism)
	}
}

func TestParallelSort(t *testing.T) {
	for parallelism := 1; parallelism <= 10; parallelism++ {
		ints := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		runParallelSortTest(ints, parallelism, t)
	}
}
