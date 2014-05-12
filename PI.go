// PI
package main

import (
	"flag"
	"fmt"
	"math/rand"
	"runtime"
)

func main() {
	nThrow := flag.Int("n", 1e6, "number of throws")
	nCPU := flag.Int("cpu", 4, "number of CPUs to use")
	flag.Parse()
	runtime.GOMAXPROCS(*nCPU) // Set number of OS threads to use.
	parts := make(chan int)   // Channel to collect partial results.
	// Kick off parallel tasks.
	for i := 0; i < *nCPU; i++ {
		go func(me int) {
			// Create local PRNG to avoid contention.
			r := rand.New(rand.NewSource(int64(me)))
			n := *nThrow / *nCPU
			hits := 0
			// Do the throws.
			for it := 0; it < n; it++ {
				x := r.Float64()
				y := r.Float64()
				if x*x+y*y < 1 {
					hits++
				}
			}
			parts <- hits // Send the result back.
		}(i)
	}
	// Aggregate partial results.
	hits := 0
	for i := 0; i < *nCPU; i++ {
		hits += <-parts
	}
	pi := 4 * float64(hits) / float64(*nThrow)
	fmt.Printf("PI = %g\n", pi)
}
