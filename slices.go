/*
* Run this example from terminal by
* using this command - go run slices.go
*
* Output:
	[2 3 5 7 11]
* Document
	The type []T is a slice with elements of type T.

	This expression creates a slice of the first five elements of the array s:

	s[0:5]
*/

package main

import (
	"fmt"
)

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[0:5]
	fmt.Println(s)
}
