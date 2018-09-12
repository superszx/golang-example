package Util

//type Enumable []struct{
//}

type Enumable struct {
	data interface{}
}

//func (source Enumable) Where(filter func(interface{}) bool) Enumable {
//	result := Enumable{make([]interface{}, 0)}
//
//	//test := []string{}
//
//	for _, item := range result.data {
//		if filter(item) {
//			//test = append(test,"asdf")
//			result.data = append(result.data, item)
//		}
//	}
//
//	return result
//}

func New(list []interface{}) Enumable {
	return Enumable{list}
}
