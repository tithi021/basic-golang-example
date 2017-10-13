/*
* Run this example from terminal by
* using this command - go run constants.go
*
* Output:
	constant
	6e+14
	600000000000000
	0.17783120151825887
*/

package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
