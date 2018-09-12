package main

import "fmt"

func mainhello() {
	// fmt.Println("test golang fmt.Println")
	fmt.Println("test golang fmt.Println")

	const (
		a       = iota
		b, c, d = iota, iota, iota
		e       = iota
	)
	fmt.Println(a, b, c, d, e)

	fmt.Printf("output = %f\n", 10.0/3)

	n1 := 52
	n2 := 40
	n3 := 2
	var avg float64
	avg = float64(n1+n2) / float64(n3)
	fmt.Printf("avg = %f\n", avg)

	fmt.Println("input a number!")
	fmt.Scanf("%d\n", &n1)
	switch n1 {
	case 1:
		fmt.Println("yi")
		//fallthrough
		//break
	case 2:
		fmt.Println("er")
		//fallthrough
	default:
		fmt.Println("san")
	}
}

// =========================
func main1() {
	i := f1()
	fmt.Println(i)
}

func f1() (r int) {
	defer func() {
		r++
	}()
	r = 0
	return
}

// ===============================
func double(x int) int {
	return x * 2
}

func triple(x int) (r int) {
	defer func() {
		r += x
	}()

	return double(x)
}

func mainhello01() {
	fmt.Println(triple(3))
}
