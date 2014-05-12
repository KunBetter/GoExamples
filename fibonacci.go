// fibonacci
package main

import (
	"fmt"
	"math"
)

//函数也是值
func func_is_value() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(3, 4))
}

//闭包
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func bibao() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(pos(i), neg(-2*i))
	}
}

//函数也是值
//fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	x := 0
	y := 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

func fib() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
