package TcpCommunication

import (
	"testing"
	"time"
)

func TestTcpCommunication(t *testing.T) {
	go func() {
		server := Server{"localhost", "8888"}
		server.Init()
	}()
	for i := 0; i < 10; i++ {
		go func() {
			client := Client{"localhost", "8888"}
			client.Init()
		}()
	}
	for {
		time.Sleep(1 * time.Second)
	}
}
