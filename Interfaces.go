// Interfaces
package main

import (
	"fmt"
	"math"
)

//Interfaces are named collections of method signatures.
type geometry interface {
	area() float64
	perim() float64
}

type square struct {
	width, height float64
}

type circle struct {
	radius float64
}

/*
For our example we will implement this interface
on square and circle types.

To implement an interface in Go,
we just need to implement all(NOTICE::all!) the methods
in the interface.
Here we implement geometry on squares.
*/
func (s square) area() float64 {
	return s.width * s.height
}

func (s square) perim() float64 {
	return 2*s.width + 2*s.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

/*
If a variable has an interface type,
then we can call methods that are in the named
interface. Here is a generic measure function
taking advantage of this to work on any geometry.
*/
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

/*
The circle and square struct types both
implement the geometry interface
so we can use instances of these structs
as arguments to measure.
*/
func interfaces_() {
	s := square{3, 4}
	c := circle{5}

	measure(&s)
	measure(c)
}
