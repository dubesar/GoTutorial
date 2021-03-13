package main

import (
	"fmt"
)

func main() {

	//array of fixed size
	var a [5]int
	a[2] = 7
	fmt.Println(a)

	//shorthand
	b := [5]int{5, 4, 3, 2, 1}
	fmt.Println(b)

	//dynamic array
	c := []int{5, 4, 3}
	c = append(c, 13)
	fmt.Println(c)
}
