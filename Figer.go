package main

import "fmt"

type Figer struct{}

func (source Figer) Test() {
	fmt.Printf("Figer trigger Test type: %T\n", source)
}
