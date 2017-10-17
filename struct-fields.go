/*
* Run this example from terminal by
* using this command - go run struct-fields.go
*
* Output:

	4
	{4 2}

* Document
	Struct fields are accessed using a dot.
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
	v := Vertex{1,2}
	v.X = 4;
	fmt.Println(v.X)
	fmt.Println(v)
}
