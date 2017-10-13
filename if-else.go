/*
* Run this example from terminal by
* using this command - go run if-else.go
*
* Output:
	7 is Odd
	8 is divisible by 4
	9 has 1 digit
*/

package main

import (
	"fmt"
)

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is Even")
	} else {
		fmt.Println("7 is Odd")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
