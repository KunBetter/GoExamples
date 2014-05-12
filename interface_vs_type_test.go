//interface_vs_type_test
package main

import (
	"sync/atomic"
	"testing"
)

func sumI(nums ...interface{}) interface{} {
	var s uint64
	for _, n := range nums {
		s += uint64(n.(int))
	}
	return s
}

func sumT(nums ...int) uint64 {
	var s uint64
	for _, n := range nums {
		s += uint64(n)
	}
	return s
}

func BenchmarkInterface(b *testing.B) {
	var s uint64
	for i := 0; i < b.N; i++ {
		s += sumI(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		s += sumI(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		s += sumI(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		s += sumI(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
	}
}

func BenchmarkType(b *testing.B) {
	var s uint64
	for i := 0; i < b.N; i++ {
		s += sumT(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s += sumT(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s += sumT(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		s += sumT(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	}
}
func BenchmarkInterfaceGoroutines(b *testing.B) {
	var s uint64
	ch, done := make(chan uint64), make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			ch <- sumI(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		}
	}
	l := func() {
		for i := 0; i < b.N; i++ {
			s += <-ch
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
		go l()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkTypeGoroutines(b *testing.B) {
	var s uint64
	ch, done := make(chan uint64), make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			ch <- sumT(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		}
	}
	l := func() {
		for i := 0; i < b.N; i++ {
			s += <-ch
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
		go l()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkInterfaceGoroutinesBuffered(b *testing.B) {
	var s uint64
	ch, done := make(chan uint64, b.N/4), make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			ch <- sumI(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64)
		}
	}
	l := func() {
		for i := 0; i < b.N; i++ {
			s += <-ch
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
		go l()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkTypeGoroutinesBuffered(b *testing.B) {
	var s uint64
	ch, done := make(chan uint64, b.N/4), make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			ch <- sumT(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		}
	}
	l := func() {
		for i := 0; i < b.N; i++ {
			s += <-ch
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
		go l()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkInterfaceGoroutinesAtomic(b *testing.B) {
	var s uint64
	done := make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			atomic.AddUint64(&s, sumI(1, 2, 3, 4, 5, 6, 7, 8, 9, 10).(uint64))
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}

func BenchmarkTypeGoroutinesAtomic(b *testing.B) {
	var s uint64
	done := make(chan bool, 4)
	f := func() {
		for i := 0; i < b.N; i++ {
			atomic.AddUint64(&s, sumT(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
		}
		done <- true
	}

	for i := 0; i < 4; i++ {
		go f()
	}

	for i := 0; i < 4; i++ {
		<-done
	}
}
