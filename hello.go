package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//while loops
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//looping arrays
	arr := []string{"a", "b", "c"}
	for index, value := range arr {
		fmt.Println("index", index, "value", value)
	}

	//same goes for map just replace index with key
}
