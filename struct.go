/*
* Run this example from terminal by
* using this command - go run struct.go
*
* Output:

	{1 2}

* Document
	A struct is a collection of fields.
	(And a type declaration does what you'd expect.) 
*/

package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
