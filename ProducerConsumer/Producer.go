package ProducerConsumer

import (
	"fmt"
	"time"
)

var PerProduceCount int = 10000

func Produce(out chan<- int, routinId int) {
	for i := 0; i < PerProduceCount; i++ {
		select {
		case out <- 1:
			fmt.Println("producer [", routinId, "] product one,current has produce ", i)
		case <-time.NewTimer(1 * time.Second).C:
			goto LoopEnd
		}
	}
LoopEnd:
	fmt.Println("producer ", routinId, " exit")
}

func Consume(in <-chan int, routinId int) {
	hasConsume := 1
	for {
		select {
		case <-in:
			fmt.Println("consumer [", routinId, "] consume one,current has consume ", hasConsume)
			hasConsume++
		case <-time.NewTimer(1 * time.Second).C:
			goto LoopEnd
		}
	}
LoopEnd:
	fmt.Println("consumer ", routinId, " exit ")
}
