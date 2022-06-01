package basics

import (
	"fmt"
	"sync"
)

var n = 1

func MutexExp() {

	var wg sync.WaitGroup
	var mt sync.Mutex

	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go incmnt(&wg, &mt)
	}

	wg.Wait()

	fmt.Println("final value of n:", n)
}

func incmnt(wg *sync.WaitGroup, mt *sync.Mutex) {

	//only one Goroutine is allowed to execute this piece of code at any point in time
	mt.Lock()
	n = n + 1
	mt.Unlock()

	wg.Done()
}
