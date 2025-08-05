package synchronization

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
func adWork(idx int) {
	time.Sleep(time.Second * time.Duration(rand.Int()))
	fmt.Printf("[GOROUTİNE %d]\n",idx)
}
func WithWaitGroup(){
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(idx int){
			defer wg.Done() 
			adWork(idx)
		}(i)
	}
	wg.Wait()
	
	fmt.Printf("[WaitGroups] all goroutines finished.")
}
func WithOutWaitGroup(){

	for i := 0; i < 10; i++ {
		go adWork(i)
	}
	fmt.Printf("[WitoutWaitGroups] finished.")

}