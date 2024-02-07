package grokkingalgorithms

var Count = 0

func QuickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	} else {
		pivot := nums[0]
		var less = []int{}
		var greater = []int{}

		for _, num := range nums[1:] {
			Count++
			if pivot > num {
				less = append(less, num)
			} else {
				greater = append(greater, num)
			}
		}

		less = append(QuickSort(less), pivot)
		greater = QuickSort(greater)
		return append(less, greater...)
	}
}
