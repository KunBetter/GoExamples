// Tickers
package main

import (
	"fmt"
	"time"
)

func Tickers_() {
	/*
		tickers are for when you want to
		do something repeatedly at regular intervals.

		Tickers use a similar mechanism to timers:
		a channel that is sent values.
		Here we will use the range builtin on the channel
		to iterate over the values as they arrive every 500ms.
	*/
	ticker := time.NewTicker(time.Microsecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	/*
		Tickers can be stopped like timers.
		Once a ticker is stopped it would not receive
		any more values on its channel.
		We will stop ours after 1500ms.
	*/
	time.Sleep(time.Microsecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped.")
}
