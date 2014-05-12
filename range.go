// range
package main

import (
	"fmt"
)

//range iterates over of elements in a variety of data structures.
func range_() {
	nums := []int{2, 3, 4}
	sum := 0
	//Here we use range to sum the numbers in a slice. Arrays work like this too.
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	//range on arrays and slices provides both the index and value for each entry.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	//range on map iterates over key/value pairs.
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range on strings iterates over Unicode code points.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
