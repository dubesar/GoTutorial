# GoTutorial
Learn Go. 

This repository contains code samples to help you learn Golang easily.
Give a ⭐️ to the repo and contribute

# Snippets
Some Short snippets for revision. Look in the files inside repo for full code and better examples.
## Basic I/O
- Input
``` go
import "fmt"
...
fmt.Scanln(&variable)
```
- Output
``` go
import "fmt"
...

fmt.Println("Hello Gophers!")
fmt.Println(variable)
```

## Variables
``` go
// Verbose way to declare variables
var x int = 5
var y int = 10

// Shortcut way to declare variables
x1 := 5
y1 := 5
sum := x1 + y1

// Printing a variable
fmt.Println(x)
fmt.Println(y)
fmt.Println(x, y)
fmt.Println("The sum of x1+y1 is: ", sum)
```

## Conditionals
``` go
if x > 4 {
	fmt.Println("x is more than 4")
} else if x == 4 {
	fmt.Println("x is equal to 4")
} else {
    fmt.Println("x is less than 4")
}
```

## Looping
``` go
for i := 0; i < 5; i++ {
	fmt.Println("Hey! ", i)
}
/* Output:
Hey!  0
Hey!  1
Hey!  2
Hey!  3
Hey!  4 */
for i := 0; i < 5; i++ {
	sum += i
}
fmt.Println(sum)
//Output: 10
```

## Arrays
``` go
var a4 [4]int           // An array of 4 ints, initialized to all 0.
a5 := [...]int{3, 1, 5, 10, 100} // An array initialized with a fixed size of five

//Initialize a Dynamic Array
arr := []int{}
arr = append(arr, 2, 3, 1, 100, 5)

// Printing the array
fmt.Println(arr)

//Print a slice of Array
fmt.Println(arr[0:2])

//Making a fixed size Array of size 4 with all Initialized by 0
new_arr := make([]int, 4)
new_arr = append(new_arr, 5)

// Sort an Array
sort.Ints(arr)
fmt.Println(arr)
//Remove an Element by Index
fmt.Println(append(arr[:1], arr[2:len(arr)-1]...))

// Remove an Element by Value
fmt.Println(indexOf(arr, 3))
fmt.Println(append(arr[:indexOf(arr, 3)], arr[indexOf(arr, 3)+1:len(arr)-1]...))
```

## Maps
Also known as a Dictionary or a Hashmap in other Languages.
``` go
currency := map[string]string{
	"AUD": "Australia Dollar",
	"GBP": "Great Britain Pound",
	"JPY": "Japan Yen",
	"CHF": "Switzerland Franc",
}

//Adding a Entry to the Map:
currency["USD"] = "USA Dollar"

//Remove a Entry from the Map:
delete(currency, "GBP")

//Replacing one Entry with Another:
currency["AUD"] = "New Zealand Dollar"

//Looping through the Map:
for key, value := range currency {
	fmt.Printf("%v might be equal to: %v\n", key, value)
}
// Printing the Map
fmt.Println("Currency Map: ", currency)
```

## Functions
``` go
func function_name(arguement arguement_type, ...) return_type {
... //Code of function
}

// Function to add 2 integers
func sum(x int, y int) int {
	return x + y
}

// Function to get the square root of a number.
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}
```

## Pointers
Special variables that can store memory address.
``` go
// Declare a pointer
var var_name *type  // The "*" here makes the variable a pointer

// Initialize a pointer with a address of another variable
var_name = &variable

// Dereference a pointer
fmt.Println("The deferenced value of pointer is: ", *var_name*)
```
