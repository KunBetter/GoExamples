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

type MinHeap struct {
	mhLen, length int
	mhA           []int32
}

func NewMinHeap(size int) *MinHeap {
	return &MinHeap{
		mhLen:  0,
		length: size,
		mhA:    make([]int32, size+1),
	}
}

func (mh *MinHeap) push(e int32) {
	if mh.mhLen == mh.length {
		if e <= mh.mhA[1] {
			return
		}
		mh.pop()
	}
	mh.mhLen++
	if mh.mhLen == 1 {
		mh.mhA[mh.mhLen] = e
		return
	}
	i := mh.mhLen
	for mh.mhA[i/2] > e {
		mh.mhA[i] = mh.mhA[i/2]
		i = i / 2
	}
	mh.mhA[i] = e
}

func (mh *MinHeap) pop() (e int32) {
	if mh.mhLen <= 0 {
		e = -1
		return
	}
	e = mh.mhA[1]
	lastE := mh.mhA[mh.mhLen]
	mh.mhA[mh.mhLen] = 0
	mh.mhLen--
	if mh.mhLen == 0 {
		return
	}
	i := 1
	for i*2 <= mh.mhLen {
		i = i * 2
		//与子节点中较小的比较，参照最小堆的定义
		/*
			子节点必须比父节点大，不然只比较一个，
			就会出现父节点比一个子节点小，
			但是比另一个大的情况
		*/
		if i+1 <= mh.mhLen && mh.mhA[i] > mh.mhA[i+1] {
			i++
		}
		if lastE > mh.mhA[i] {
			mh.mhA[i/2] = mh.mhA[i]
		} else {
			i = i / 2
			break
		}
	}
	mh.mhA[i] = lastE
	return
}

func (mh *MinHeap) lazyGrow() {
	if mh.full() {
		mh.grow()
	}
}

func (mh *MinHeap) full() bool {
	return mh.mhLen == mh.length
}

func (mh *MinHeap) grow() {
	big := make([]int32, mh.length*2+1)
	for i := 1; i <= mh.mhLen; i++ {
		big[i] = mh.mhA[i]
	}
	mh.mhA = big
	mh.length *= 2
}

func (mh *MinHeap) test() {
	v := []int32{188, 17, 22, 21, 99, 30, 19, 20, 200, 180, 1000}
	hv := []int32{}
	for i := 0; i < len(v); i++ {
		mh.push(v[i])
	}
	fmt.Println(mh.mhA)
	mh.lazyGrow()
	fmt.Println(mh.mhA)
	fmt.Println(mh)
	for mh.mhLen > 0 {
		hv = append(hv, mh.pop())
	}
	fmt.Println(hv)
}

func main() {
	mh := NewMinHeap(8)
	mh.test()
}
