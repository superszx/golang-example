package DeadLockCase

import (
	"fmt"
	"time"
)

func Case1() { //no write end lead to deadlock
	readCh := make(<-chan bool)
	any := <-readCh
	fmt.Println(any)
}

func Case2() { //block by read end ,lead to write end can not be run , lead to deadlock
	readCh := make(chan bool)
	//go func (){
	//	readCh <- true
	//}()
	any := <-readCh
	fmt.Println(any)
	go func() {
		readCh <- true
	}()
}

func Case3() { //the sync order with read write channel ,run correctly
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func() {
		for {
			select {
			case <-ch1:
				{
					fmt.Println("sub routin read 1 ")
					<-ch2
					fmt.Println("sub routin read 2")
				}
			case <-time.NewTimer(1 * time.Millisecond).C:
				fmt.Println("time out break one loop")
				break
				//default:
				//	 println("can not read ,busy now")
			}
		}
	}()

	for {
		select {
		case ch1 <- true:
			{
				println("main routin write 1")
				ch2 <- true
				println("main routin write 2")
			}
		case <-time.NewTimer(1 * time.Millisecond).C:
			fmt.Println("time out break one loop")
			break
			//default:
			//	println("can not write ,busy now")
		}
	}
}

func Case4() { // un property order lead to deadlock
	ch1 := make(chan bool)
	ch2 := make(chan bool)
	go func() {
		for {
			select {
			case <-ch1:
				{
					fmt.Println("sub routin read 1 ")
					<-ch2
					fmt.Println("sub routin read 2")
				}
			}
		}
	}()

	for {
		select {
		case ch2 <- true:
			{
				println("main routin write 2")
				ch1 <- true
				println("main routin write 1")
			}
		}
	}
}
