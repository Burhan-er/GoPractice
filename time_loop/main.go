package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 10000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)
	fmt.Printf("total time without preallocation: %v", timeLoop(testSlice, n))
	fmt.Printf("total time with preallocation: %v", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now()
	for len(slice) < n {
		slice = append(slice, 1)
	}
	var t1 = time.Now()
	duration := t1.Sub(t0)
	return duration
}
