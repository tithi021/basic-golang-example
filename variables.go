/*
* Run this example from terminal by
* using this command - go run variables.go
*
* Output:
	Initial string
	1 2
	true
	0
	short
*/

package main

import (
	"fmt"
)

func main() {

	// String type variable
	var a string = "Initial string"
	fmt.Println(a)

	// Integer type variables
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Boolean type variable
	var d = true
	fmt.Println(d)

	// Integer
	var e int
	fmt.Println(e)

	// := syntax is shorthand for declaring and initializing a variable
	f := "short"
	fmt.Println(f)
}
