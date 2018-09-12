package ParallelDemo

import (
	"testing"
	"time"
)

func TestIncreaseIntWithoutLock(t *testing.T) {
	IncreaseIntWithoutLock()
}

func TestIncreaseIntWithLock(t *testing.T) {
	IncreaseIntWithLock()
}

//--- PASS: TestIncreaseIntWithoutLock (1.50s)
//PASS: TestIncreaseIntWithLock (132.96s)

func Test(t *testing.T) {
	cond.L = mutex
	for i := 0; i < producerCount; i++ {
		go ProducerHandler(i)
	}

	for i := 0; i < consumerCount; i++ {
		go ConsumerHandler(i)
	}
	for {
		time.Sleep(1 * time.Second)
	}
}
