// WorkerPools
package main

import (
	"fmt"
	"time"
)

//implement a worker pool using goroutines and channels.
/*
Here is the worker,
of which we will run several concurrent instances.
These workers will receive work on the jobs channel
and send the corresponding results on results.
We will sleep a second per job to simulate an expensive task.
*/
func worker_(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func WorkerPools_() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	/*
		This starts up 3 workers,
		initially blocked because there are no jobs yet.
	*/
	for w := 1; w <= 3; w++ {
		go worker_(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= 9; a++ {
		<-results
	}
	/*
		Our running program shows the 9 jobs being executed
		by various workers. The program only takes about 3 seconds
		despite doing about 9 seconds of total work
		because there are 3 workers operating concurrently.
	*/
}
