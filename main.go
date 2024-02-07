package main

import (
	"fmt"
	algorithms "grokking_algorithms/grokking_algorithms"
)

func main() {
	// fmt.Println(algorithms.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	// fmt.Println(algorithms.LinearSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))

	// fmt.Println(algorithms.SelectionSort([]int{4, 2, 6, 8, 21, 12, 10, 25, 29, 32, 61, 98, 100}))
	fmt.Println(algorithms.QuickSort([]int{4, 2, 6, 8, 21, 12, 10, 25, 29, 32, 61, 98, 100}))
}
