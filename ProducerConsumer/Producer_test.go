package ProducerConsumer

import (
	"testing"
)

func TestConsume(t *testing.T) {
	channel := make(chan int, 0)
	for i := 1; i <= 5; i++ {
		go Produce(channel, i)
	}

	Consume(channel, 1)
}
