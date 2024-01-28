package main

import "fmt"

func findMe(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		// Calculate middle index
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			// If target is greater, focus on the right half
			left = mid + 1
		} else {
			// If target is less, focus on the left half
			right = mid - 1
		}
	}
	// if tartget not fount return -1
	return -1
}

func main() {
	fmt.Println(findMe([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 8)) //resp 7
	fmt.Println(findMe([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 4)) //resp 3
}
