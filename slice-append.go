/*
* Run this example from terminal by
* using this command - go run slice-append.go
*
* Output:
	Len=0 cap=0 []
	Len=1 cap=1 [0]
	Len=2 cap=2 [0 1]
	Len=5 cap=6 [0 1 2 3 4]
* Document
	It is common to append new elements to a slice, and so Go provides a built-in append function
	func append(s []T, vs ...T) []T
*/

package main

import (
	"fmt"
)

func main() {
	var s []int
	printSlice(s)

	// Append works on nil slices
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("Len=%d cap=%d %v\n", len(s), cap(s), s)
}
