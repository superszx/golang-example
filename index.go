package main

import (
	"fmt"
	"golang-example/service"
)

func mainindex() {
	service1.Login("myname", "mypwd")

	a, b := 3, 4

	fmt.Println("before swap", a, b)

	service1.Util{}.Swap(&a, &b)

	fmt.Println("after swap", a, b)
}
