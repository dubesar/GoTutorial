package main

import "fmt"

func linearsearch(array []int, key int) bool {
	for _, item := range array {
		if item == key {
			return true
		}
	}
	return false
}

func main() {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(linearsearch(items, 8))  // true
	fmt.Println(linearsearch(items, 23)) // false
}
