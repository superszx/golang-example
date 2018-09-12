package ParallelDemo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var LoopCount int = 1000000000
var LoopRountinCount int = 1

var producerCount = 5
var consumerCount = 1

func IncreaseIntWithoutLock() {
	wg := sync.WaitGroup{}
	var cur int = 0
	for i := 0; i < LoopRountinCount; i++ {
		wg.Add(1)
		go func(i int) {
			WaitHandler := func(i int) func() {
				return func() {
					wg.Done()
					fmt.Println("routine ", i, " completed")
				}
			}(i) //闭包
			//fmt.Println(i)
			defer WaitHandler()
			for j := 0; j < LoopCount; j++ {
				cur++
			}
		}(i) //define as a closure,and execute after loop end
	}
	fmt.Println("waiting for complete...")
	wg.Wait()
	fmt.Println("cur = ", cur)
	fmt.Println("done...")
}

func IncreaseIntWithLock() {
	wg := sync.WaitGroup{}
	cur := 0
	increamentMutex := sync.Mutex{}
	for i := 0; i < LoopRountinCount; i++ {
		wg.Add(1)
		go func(i int) {
			defer func() {
				wg.Done()
				fmt.Println("routine ", i, " completed")
			}()
			for j := 0; j < LoopCount; j++ { //with mutex
				increamentMutex.Lock()
				cur++
				increamentMutex.Unlock()
			}
		}(i)
	}
	fmt.Println("waiting for complete...")
	wg.Wait()
	fmt.Println("cur = ", cur)
	fmt.Println("done...")
}

//var ch = make(chan int)
var ch = make(chan int, 1000)
var cond = new(sync.Cond)

// in no cache channel , mutex will block read routine to read ,but write routine need it online
// in with cache channel , mutex block read routine to read, but due to the capability of channel ,write channel also be blocked
var mutex = new(sync.Mutex)

func ProducerHandler(i int) {
	rand.Seed(time.Now().UTC().UnixNano())

	for {
		//mutex.Lock()
		cond.L.Lock()
		for len(ch) >= 5 {
			cond.Wait()
		}
		sed := rand.Intn(10000)
		ch <- sed
		fmt.Println("producer ", i, " product ", sed)
		cond.L.Unlock()
		cond.Broadcast()

		//mutex.Unlock()
	}
}

func ConsumerHandler(i int) {
	for {

		//mutex.Lock()
		cond.L.Lock()
		for len(ch) <= 0 {
			cond.Wait()
		}
		rec := <-ch
		fmt.Println("consumer ", i, "consume ", rec)
		cond.L.Unlock()
		cond.Broadcast()
		//mutex.Unlock()
	}

}
