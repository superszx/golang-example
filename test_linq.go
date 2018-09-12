package main

//
//import "fmt"
//import (
//	"golang-example/Util"
//	. "golang-example/Util"
//	"strings"
//)
//
//func maintestlinq() {
//	fmt.Println("*************linq test ****************")
//	// l  := []struct {
//	//	name string
//	//	value string
//	//}{{"testname","testvalue"},{"1","1"},{"2","2"}}
//
//	var list = Enumable{make([]Product, 0)}
//	//var dataList Enumable = Util.New([]Product)
//	//var l Emumable = Emumable{{name:"testname",value:"testvalue"},{"1","1"},{"2","2"}}
//	result := dataList.Where(func(item interface{}) bool {
//		p := new(Product)
//		return p.ProNo == "8806"
//	})
//
//	fmt.Println(result)
//
//	fmt.Println("*******************test list***********************")
//
//	var lis StringList = []string{"asdf", "dddd", "ccccc"}
//
//	fmt.Println(lis)
//
//	lis = lis.Up()
//
//	fmt.Println(lis)
//	fmt.Println("*******************list contaner***********************")
//
//	fmt.Printf("%T", list.New())
//
//	// ll := list.New()
//
//}
//
//// struct Product ... ProNo shangpingProName
//// dfa
//type Product struct {
//	// sdfds
//	ProNo string
//	// sdf
//	ProName string
//}
//
//type StringList []string
//
//// ...Up 转换大小写
//func (source StringList) Up() []string {
//	result := []string{}
//	for _, item := range source {
//		result = append(result, strings.ToUpper(item))
//	}
//	return result
//}
//
////func ConvertToEnumable( source []interface{}) []Emumable{
////	result := []Emumable
////	for _,item := range source{
////		result = append(result,Emumable(item))
////	}
////}
