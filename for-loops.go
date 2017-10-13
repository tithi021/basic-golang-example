/*
* Run this example from terminal by
* using this command - go run for-loops.go
*
* Output:
	1
	2
	3
	7
	8
	9
	loop
	1
	3
	5
*/

package main

import (
	"fmt"
)

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
