// Slices
package main

import (
	"fmt"
)

func slices_() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("get:", s)

	fmt.Println("s's addr:", &s)
	s = append(s, "d")
	fmt.Println("s's addr:", &s)
	s = append(s, "e", "f")
	fmt.Println("s's addr:", &s)
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	/*
		slice通过array[i:j]来获取，其中i为数组的开始位置,
		j为结束位置,点不包括j,它的长度为j-i,
		如果不指定i,默认从开头开始,不指定j,
		则直到数组的最后一位.
		注意: slice的引用是内存地址,所以当改变其中的元素的值时,
		其他的所有指向该slice都会改变该值.
		len()获取slice的长度,cap()获取slice的最大容量,
		slice的容量为提取数组的开始位置到数组结束位置.
	*/
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D:", twoD)
}

func slice_copy() {
	a := []int{1, 2, 3, 4}
	b := []int{5, 6, 7, 8}
	c := []int{9, 10, 11, 12}

	alen := len(a)
	blen := len(b)
	clen := len(c)

	d := make([]int, alen+blen+clen+1)

	copy(d[1:], a)
	copy(d[1+alen:], b)
	copy(d[1+alen+blen:], c)

	fmt.Println(d)
}

func main1() {
	//slices_()
	slice_copy()
}
