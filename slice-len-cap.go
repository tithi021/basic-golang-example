/*
* Run this example from terminal by
* using this command - go run slice-len-cap.go
*
* Output:
	len=6 cap=6 [2 3 5 7 11 13]
	len=0 cap=6 []
	len=4 cap=6 [2 3 5 7]
	len=2 cap=4 [5 7]
* Document
	A slice has both a length and a capacity.

	The length of a slice is the number of elements it contains.

	The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

	The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).
*/

package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	PrintSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	PrintSlice(s)

	// Extend It's Length
	s = s[:4]
	PrintSlice(s)

	// Drop It's first two values
	s = s[2:]
	PrintSlice(s)
}

func PrintSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
