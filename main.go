package main

import (
	"fmt"
	algorithms "grokking_algorithms/grokking_algorithms"
)

func main() {
	fmt.Println(algorithms.BinarySearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
	fmt.Println(algorithms.LinearSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
