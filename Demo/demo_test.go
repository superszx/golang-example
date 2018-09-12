package Demo

import (
	"fmt"
	"testing"
)

func TestRemoveEmptyStringNoReturn0(t *testing.T) {
	s := []string{"sdfa", "", "sdfa", ""}
	RemoveEmptyStringNoReturn0(s)
	fmt.Println(s)

	// output
	//[sdfa  sdfa ]
}

func TestRemoveEmptyStringNoReturn(t *testing.T) {
	s := []string{"sdfa", "", "sdfa", ""}
	RemoveEmptyStringNoReturn(&s)
	fmt.Println(s)
	//output
	//[sdfa sdfa]
}

func TestRemoveEmtyStringWithReturn(t *testing.T) {
	s := []string{"sdfa", "", "sdfa", ""}
	tmp := RemoveEmptyStringWithReturn(s)
	fmt.Println(tmp)
	//output
	//[sdfa sdfa]
}

func TestDistincStringList(t *testing.T) {
	s := []string{"aaa", "aaa", "bbb", "bbb"}
	s1 := DistincStringList(s)
	fmt.Println(s1)
	// output
	// [aaa bbb]
}
