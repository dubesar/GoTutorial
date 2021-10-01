package main

func MergeSort(list []int) []int {
	return sort(list)
}

func sort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	left, right := split(list)
	left = sort(left)
	right = sort(right)
	list = merge(left, right)
	return list

}
func split(list []int) ([]int, []int) {
	return list[0 : len(list)/2], list[len(list)/2:]
}

func merge(left []int, right []int) []int {
	sorted := make([]int, (len(left) + len(right)))

	leftIndex := 0
	rightIndex := 0

	for i := 0; i <= (len(sorted) - 1); i++ {
		if leftIndex >= len(left) {
			sorted[i] = right[rightIndex]
			rightIndex++
			continue
		} else if rightIndex >= len(right) {
			sorted[i] = left[leftIndex]
			leftIndex++
			continue
		}

		if left[leftIndex] <= right[rightIndex] {
			sorted[i] = left[leftIndex]
			leftIndex++
		} else if right[rightIndex] <= left[leftIndex] {
			sorted[i] = right[rightIndex]
			rightIndex++
		}

	}

	return sorted
}
