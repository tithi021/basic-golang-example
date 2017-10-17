/*
* Run this example from terminal by
* using this command - go run maps.go
*
* Output:
	map: map[k2:14 k1:7]
	v1:  7
	len: 2
	map: map[k1:7]
	prs: false
	map: map[foo:1 bar:2]
* Document
	To create an empty map, use the builtin make: make(map[key-type]val-type).
*/

package main

import (
	"fmt"
)

func main() {

	// create map
	// make(map[key-type]val-type).
	m := make(map[string]int)

	m["k1"] = 7;
	m["k2"] = 14

	fmt.Println("map:" , m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n:= map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
