/*
* Run this example from terminal by
* using this command - go run slice-bounds.go
*
* Output:
	[3 5 7]
	[3 5]
	[5]
* Documents
	When slicing, you may omit the high or low bounds to use their defaults instead.

	The default is zero for the low bound and the length of the slice for the high bound.

	For the array

	var a [10]int

	these slice expressions are equivalent:

	a[0:10]
	a[:10]
	a[0:]
	a[:]
*/

package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
