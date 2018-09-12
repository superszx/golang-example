package main

import "fmt"

type Cracker struct {
}

func (source Cracker) Test() {
	fmt.Printf("creacker trigger Test , type is %T\n", source)
}
