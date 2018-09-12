package Linq

import (
	"fmt"
	"testing"
)

func TestEnumable_NotEmpty(t *testing.T) {
	s := New([]string{"dfd", "sdfas", "", "sdfasd", "", "asdfasdf"}).NotEmpty()

	fmt.Printf("%v", s)
	// output :
	// [dfd sdfas sdfasd asdfasdf]

	//s1 := New([]int {1,2,1,0,1,0,}).NotEmpty() //panic

	//fmt.Printf("%v",s1)

}
