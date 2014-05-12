// UnionFindSet
package main

import (
	"fmt"
)

var (
	Father = make([]int, 1000, 1000)
)

func init_() {
	for i := 0; i < len(Father); i++ {
		Father[i] = i
	}
}

func getFather(x int) int {
	//路径压缩
	if Father[x] != x {
		Father[x] = getFather(Father[x])
	}
	return Father[x]
}

func union(x, y int) {
	Father[getFather(x)] = getFather(y)
}

func same(x, y int) bool {
	if getFather(x) != getFather(y) {
		return false
	}
	return true
}

func main() {
	init_()
	union(2, 3)
	union(4, 5)
	union(6, 7)
	union(8, 9)
	union(2, 9)
	union(3, 7)
	fmt.Println(same(4, 8))
	fmt.Println(same(5, 6))
	fmt.Println(same(6, 8))
	fmt.Println(same(2, 6))
}
