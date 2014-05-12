// Goroutine
package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func goruntine_() {
	f("direct")

	go f("goruntine")

	//anonymous function
	go func(msg string) {
		fmt.Println(msg)
		for i := 0; i < 3; i++ {
			fmt.Println(msg, "::", i)
		}
	}("going")

	fmt.Println("Hello World!")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done!")
}
