package main

import (
	"fmt"
)

func main() {
	vertices := make(map[string]int)
	vertices["triangle"] = 2
	vertices["square"] = 3
	vertices["dodecagon"] = 1

	fmt.Println(vertices)

	//to delete any key
	delete(vertices, "square")
	fmt.Println(vertices)
}
