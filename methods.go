// methods
package main

import (
	"fmt"
)

type rect struct {
	width, height int
}

//Go supports methods defined on struct types.
func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func methods_() {
	r := rect{10, 5}

	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	/*
		Go automatically handles conversion between values
		and pointers for method calls. You may want to use
		a pointer receiver type to avoid copying on method
		calls or to allow the method to mutate
		the receiving struct.
	*/
	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}
