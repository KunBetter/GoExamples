// AtomicCounters
package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func AtomicCounters_() {
	var ops uint64 = 0
	for i := 0; i < 50; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
			runtime.Gosched()
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
