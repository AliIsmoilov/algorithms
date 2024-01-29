package grokkingalgorithms

func LinearSearch(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return nums[i]
		}
	}
	return -1
}



