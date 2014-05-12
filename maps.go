// maps
package main

import (
	"fmt"
)

func maps_() {
	//To create an empty map, use the builtin make:
	//make(map[key-type]val-type).
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	/*
		The optional second return value when
		getting a value from a map indicates
		if the key was present in the map.
		This can be used to disambiguate between
		missing keys and keys with zero
		values like 0 or "".
	*/
	_, pr1 := m["k1"]
	fmt.Println("prs:", pr1)

	_, pr2 := m["k2"]
	fmt.Println("prs:", pr2)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
