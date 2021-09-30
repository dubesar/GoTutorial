package main

import (
	"fmt"
	"sort"
)

//searching element
//go doesnt provide any package for searching
func indexOf(data []int, ele int) int {
	for index, value := range data {
		if ele == value {
			return index
		}
	}
	return -1
}

func main() {
	//making dynamic array

	//initialize dynamic arr
	arr := []int{}
	arr = append(arr, 2, 3, 1, 100, 5)

	fmt.Println(arr)

	//print a slice of array
	fmt.Println(arr[0:2])

	//making a fixed size array of size 4 with all initialized by 0
	new_arr := make([]int, 4)
	new_arr = append(new_arr, 5)

	fmt.Println(new_arr)

	//arr operations ->
	//get an element at a particular location
	fmt.Println(arr[2])

	// sort an array
	sort.Ints(arr)
	fmt.Println(arr)

	//remove an element
	//by index
	fmt.Println(append(arr[:1], arr[2:len(arr)-1]...))

	//by value
	fmt.Println(indexOf(arr, 3))
	fmt.Println(append(arr[:indexOf(arr, 3)], arr[indexOf(arr, 3)+1:len(arr)-1]...))
}
