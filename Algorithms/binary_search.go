package main

import "fmt"
func binarySearch(target int, arr []int) bool {
	low := 0
	high := len(arr) - 1

	for low <= high{
		mid := (low + high) / 2

		if arr[mid] < target {
			low = mid + 1
		}else if arr[mid]==target{
		    return true
		}else{ 
			high = mid - 1
		}
	}
	return false
}
func main() {
	items := []int{1,2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(31, items))
}