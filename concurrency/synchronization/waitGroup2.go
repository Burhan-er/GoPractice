package synchronization

import (
	"fmt"
	"sync"
	"time"
)

func task1() {
	time.Sleep(time.Second * 5)
}
func task2() {
	time.Sleep(time.Second * 2)
}
func task3() {
	time.Sleep(time.Second * 9)
}

func WithWaitGroup2() {
	initTime := time.Now()
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		task1()
		wg.Done()
	}()
	go func() {
		task2()
		wg.Done()
	}()
	go func() {
		task3()
		wg.Done()
	}()
	wg.Wait()
	finishTime := time.Now()
	fmt.Printf("All Tasks Finished! in %d seconds",finishTime.Second()-initTime.Second())
}
