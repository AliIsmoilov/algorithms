package grokkingalgorithms

import (
	"fmt"
	"time"
)

func CountDown(num int) {
	time.Sleep(time.Second * 1)
	fmt.Println(num)
	if num <= 0 {
		return
	} else {
		CountDown(num - 1)
	}
}

// This `factorial` function
// using recurion
func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}

// Here is the `factorial` function without recursion
// for comparation
func FactorialWitoutRecursion(n int) int {
	fact := 1
	for i := 1; i <= n; i++ {
		fact = fact * i
	}

	return fact
}

// SumOfArrayRecursive returns the sum of
// array numbers
// using recursion
func SumOfArrayRecursive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + SumOfArrayRecursive(nums[1:])
}
