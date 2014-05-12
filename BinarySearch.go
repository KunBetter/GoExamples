// BinarySearch
package main

import (
	"fmt"
)

func bs(d []int, n int) int {
	l := 0
	h := len(d) - 1
	for l <= h {
		m := (l + h) / 2
		if d[m] < n {
			l = m + 1
		} else if d[m] > n {
			h = m - 1
		} else {
			return m
		}
	}
	return -1
}

func bs_repeat(d []int, n int) (low, high int) {
	/*
		对上下限的条件判断，比如是否有等号，
		可以举例判断。
		求上限时，可以用8,8,9这样的数组判断;
		求下限时，可以用8,9,9这样的数组判断.
	*/
	l := 0
	h := len(d) - 1
	for l < h {
		m := (l + h + 1) / 2
		if d[m] <= n {
			l = m
		} else {
			h = m - 1
		}
	}
	high = h
	l = 0
	h = high
	for l < h {
		m := (l + h) / 2
		if d[m] >= n {
			h = m
		} else {
			l = m + 1
		}
	}
	low = l
	return
}

func main() {
	d := []int{4, 5, 6, 9, 10, 11, 15, 17, 19}
	fmt.Println(d)
	for i := 0; i < len(d); i++ {
		fmt.Println(bs(d, d[i]))
	}
	dr := []int{4, 4, 5, 5, 9, 9, 9, 10, 11, 11}
	fmt.Println(dr)
	for i := 0; i < len(dr); i++ {
		fmt.Println(bs_repeat(dr, dr[i]))
	}
}
