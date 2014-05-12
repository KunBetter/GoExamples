// minheap
package main

import (
	"fmt"
)

/*
1、把数组的第一个节点空余。
2、按照数组下标，下标为n的节点，它的子结点下标为2*n和2*n+1;
3、插入节点时，先插入到最后，然后再调整堆。
4、删除最小节点即删除根节点，先将根节点和最后一个节点交换，再调整堆。
*/

const length = 8

var (
	minHA = make([]int, length+1)
	mhLen = 0
)

func push(e int) {
	if mhLen == length {
		if e <= minHA[1] {
			return
		}
		pop()
	}
	mhLen++
	if mhLen == 1 {
		minHA[mhLen] = e
		return
	}
	i := mhLen
	for minHA[i/2] > e {
		minHA[i] = minHA[i/2]
		i = i / 2
	}
	minHA[i] = e
}

func pop() (e int) {
	if mhLen <= 0 {
		e = -1
		return
	}
	e = minHA[1]
	lastE := minHA[mhLen]
	minHA[mhLen] = 0
	mhLen--
	if mhLen == 0 {
		return
	}
	i := 1
	for i*2 <= mhLen {
		i = i * 2
		//与子节点中较小的比较，参照最小堆的定义
		/*
			子节点必须比父节点大，不然只比较一个，
			就会出现父节点比一个子节点小，
			但是比另一个大的情况
		*/
		if i+1 <= mhLen && minHA[i] > minHA[i+1] {
			i++
		}
		if lastE > minHA[i] {
			minHA[i/2] = minHA[i]
		} else {
			i = i / 2
			break
		}
	}
	minHA[i] = lastE
	return
}

func main() {
	v := []int{18, 17, 22, 21, 99, 30, 19, 20, 100, 180, 1000}
	for i := 0; i < len(v); i++ {
		push(v[i])
	}
	fmt.Println(minHA)
	for mhLen > 0 {
		fmt.Println(pop())
		fmt.Println(minHA)
	}
}
