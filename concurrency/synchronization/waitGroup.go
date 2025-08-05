package synchronization

import (
	"fmt"
	"sync"
	"time"
)

func adWork(idx int) {
	time.Sleep(time.Second * 1)
	fmt.Printf("[GOROUTİNE %d]\n", idx)
}
func WithWaitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			adWork(idx)
		}(i)
	}
	wg.Wait()

	fmt.Printf("[WaitGroups] all goroutines finished.")
}
func WithOutWaitGroup() {

	for i := 0; i < 10; i++ {
		go adWork(i)
	}
	time.Sleep(time.Second * 4)
	fmt.Printf("[WitoutWaitGroups] finished.")

}
