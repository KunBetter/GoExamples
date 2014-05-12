// Timers
package main

import (
	"fmt"
	"time"
)

func timers_() {
	/*
		Timers represent a single event in the future.
		You tell the timer how long you want to wait,
		and it provides a channel that will be notified
		at that time. This timer will wait 2 seconds.
	*/
	time1 := time.NewTimer(time.Second * 2)

	/*
		The <-timer1.C blocks on the timer is channel C
		until it sends a value indicating
		that the timer expired.
	*/
	<-time1.C
	fmt.Println("Timer 1 expired.")

	time2 := time.NewTimer(time.Second)
	go func() {
		<-time2.C
		fmt.Println("Timer 2 expired.")
	}()
	stop2 := time2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped.")
	}
}
