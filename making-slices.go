/*
* Run this example from terminal by
* using this command - go run making-slices.go
*
* Output:
	a len=5 cap=5 [0 0 0 0 0]
	b len=0 cap=5 []
	c len=2 cap=5 [0 0]
	d len=3 cap=3 [0 0 0]
* Document
	Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

	The make function allocates a zeroed array and returns a slice that refers to that array:

	a := make([]int, 5)  // len(a)=5
	To specify a capacity, pass a third argument to make:

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5

	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
*/

package main

import (
	"fmt"
)

func main() {
	a := make([]int, 5)
	PrintSlice("a", a)

	b := make([]int, 0, 5)
	PrintSlice("b", b)

	c := b[:2]
	PrintSlice("c", c)

	d := c[2:5]
	PrintSlice("d", d)
}

func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
