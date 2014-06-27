// FlagOfNetherlands
// Flag Netherlands
package main

import (
	"fmt"
)

//http://blog.csdn.net/v_july_v/article/details/6211155
func FlagOfNetherlands(flags []int) {
	fLen := len(flags)
	head := 0
	tail := fLen - 1
	for cur := head; cur <= tail; cur++ {
		fmt.Println(flags)
		if flags[cur] == 0 {
			flags[cur], flags[head] = flags[head], flags[cur]
			head++
		} else if flags[cur] == 2 {
			flags[cur], flags[tail] = flags[tail], flags[cur]
			tail--
			cur--
		} else {
		}
	}
}

func main() {
	flags := []int{0, 1, 2, 1, 1, 2, 0, 2, 1, 0}
	FlagOfNetherlands(flags)
}
