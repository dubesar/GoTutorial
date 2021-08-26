package main

import (
	"fmt"
)

func main() {
	//looping arrays
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}
	//same goes for map just replace index with key
	
	//using for loops
	sum := 0
	for i := 0; i < 5; i++ {
	    sum += i
	}
	fmt.Println(sum)

	n := 25
	for sum < 100 {
	    sum += n
	}

	fmt.Println(sum)
}
