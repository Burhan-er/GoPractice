package synchronization

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

func mDoWork(balance *int, withMutex bool) {
	if withMutex {
		mutex.Lock()
		*balance++
		mutex.Unlock()
	} else {
		*balance++
	}
	time.Sleep(time.Second*1)

}

func WithMutex(){
	initTime := time.Now()
	balance := 0

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			mDoWork(&balance,true)
		}()
	}

	wg.Wait()
	
	finishTime := time.Now()
	fmt.Printf("[WithMutex] finished at %d nanoseconds. Balance:%d\n",finishTime.Nanosecond()-initTime.Nanosecond(),balance)
}
func WithoutMutex(){
	initTime := time.Now()
	balance := 0

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			mDoWork(&balance,false)
		}()
	}

	wg.Wait()
	
	finishTime := time.Now()
	fmt.Printf("[WithoutMutex] finished at %d nanoseconds. Balance:%d\n",finishTime.Nanosecond()-initTime.Nanosecond(),balance)
}
