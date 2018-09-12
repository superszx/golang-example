package service1

import "fmt"

func Login(userName, pwd string) bool {
	fmt.Println(userName, pwd)
	return true
}
