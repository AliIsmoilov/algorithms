package grokkingalgorithms

func FindSmallest(arr []int) int {
	// stores the smallest value
	smallest := arr[0]
	// Stores the index of the smallest value
	smallest_index := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_index = i
		}
	}
	return smallest_index
}

func FindBiggest(arr []int) int {
	biggest := arr[0]
	biggest_index := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] > biggest {
			biggest = arr[i]
			biggest_index = i
		}
	}
	return biggest_index
}

func SelectionSort(arr []int) []int {
	size := len(arr)
	newArr := make([]int, size)

	for i := 0; i < size; i++ {
		// finds the smallest value from arr
		smallest := FindSmallest(arr)
		// stores the smallest value to the arr
		newArr[i] = arr[smallest]
		// filteres out the used value from arr
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}
	return newArr
}
