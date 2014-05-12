// QuickSort
package main

import (
	"fmt"
)

func quickSort(array []int, start int, end int) {
	if start >= end {
		return
	}
	//以最后一个数据为基准点
	s := end
	//i为从头开始遍历时小于基准的游标
	i := start - 1
	//j为从头开始遍历时大于基准的游标
	j := start
	for j <= s {
		//添加=的意思是将最后的基准点调整到小于
		//和大于它的中间
		if array[j] <= array[s] {
			if i+1 != j {
				array[i+1], array[j] = array[j], array[i+1]
			}
			i++
		}
		j++
	}
	quickSort(array, start, i-1)
	quickSort(array, i+1, end)
}

func main() {
	v := []int{2, 8, 7, 1, 3, 5, 6, 4, 0, 9, 10, 12, 11}
	quickSort(v, 0, len(v)-1)
	fmt.Println(v)

	v = []int{2, 3, 4, 5, 6, 7, 8}
	quickSort(v, 0, len(v)-1)
	fmt.Println(v)

	v = []int{8, 7, 6, 5, 4, 3, 2}
	quickSort(v, 0, len(v)-1)
	fmt.Println(v)
}
