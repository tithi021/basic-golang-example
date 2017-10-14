/*
* Run this example from terminal by
* using this command - go run slices-pointers.go
*

* Output:
	[John Paul George Ringo]
	[John Paul] [Paul George]
	[John XXX] [XXX George]
	[John XXX George Ringo]

* Document
	A slice does not store any data, it just describes a section of an underlying array.
	Changing the elements of a slice modifies the corresponding elements of its underlying array.
	Other slices that share the same underlying array will see those changes.
*/

package main

import (
	"fmt"
)

func main() {

	// Define string type of array
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	// Print
	fmt.Println(names)

	// Slice of first 2 elements of array names
	a := names[0:2]
	// Slice of 1-3 elements of array names
	b := names[1:3]
	// Print
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
