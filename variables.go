package main

import (
	"fmt"
)

func main() {
	var x int = 5
	var y int = 10
	fmt.Println(x + y)

	//similarly
	x1 := 5
	y1 := 5
	sum := x1 + y1
	fmt.Println(sum)

	//printing
	fmt.Println("Hello, World!")

	//using variables
	a := "hello"

	//once a variable is defined it can used as
	a = "world"

	//type can be given to a variable as
	var b string

	//condition
	if a == "hello" {
		fmt.Println("a is hello")
	} else if a == "world" {
		fmt.Println("a is world")
	} else {
		fmt.Println("a is not hello")
	}

	if b == "" {
		fmt.Println("do not keep any variable unused")
	}
}
