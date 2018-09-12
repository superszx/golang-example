package main

import (
	"fmt"
	"golang-example/Linq"
)

func main() {
	fmt.Println("***************slice******************")
	arr1 := [5]int{0, 1, 2, 3, 4}
	arr2 := arr1
	s1 := arr1[0:]
	fmt.Println(arr1, arr2, s1)
	fmt.Printf("%p %p %p \n", &arr1, &arr2, &s1)
	arr1[1] = 0
	fmt.Println(arr1, arr2, s1)
	fmt.Printf("%p %p %p \n", &arr1, &arr2, &s1)

	fmt.Println("***************slice copy ******************")
	s2 := []int{6, 6, 6}
	copy(s1, s2)
	fmt.Println(s1, s2)
	fmt.Printf("%p %p \n", &s1, &s2)
	fmt.Println("***************slice cap ******************")
	s1 = append(s1, 5)
	fmt.Printf("%p \n", &s1)
	for i := 0; i < 100; i++ {
		s1 = append(s1, i)
	}

	fmt.Printf("%p \n", &s1)

	fmt.Println("***************interface******************")
	var testser ITestable = new(Cracker)
	var figer ITestable = new(Figer)
	testser.Test()
	figer.Test()

	printType(testser)
	printType(figer)
	fmt.Println("***************make new ******************")

	printType(make([]int, 3))
	printType(new([]int))
	printType(new(Cracker))

	fmt.Println("***************remove empty string ******************")
	containEmptyStringList := []string{"dfas", "", "", "sdfas", ""}
	notContainEmptyStringList := Linq.New(containEmptyStringList).NotEmpty()
	fmt.Printf("notContainEmptyStringList = %v\n", notContainEmptyStringList)
	fmt.Printf("containEmptyStringList = %v\n", containEmptyStringList)

}
