package Linq

import (
	"reflect"
)

func (source Enumable) NotEmpty() []string {
	data := reflect.ValueOf(source.source)

	res := make([]string, 0)

	switch data.Kind() {
	case reflect.Slice:
		len := data.Len()
		for i := 0; i < len; i++ {
			if data.Index(i).Kind() != reflect.String {
				panic("invaild parameter, hope type is []string but input type is " + data.Index(i).Kind().String())
			}
			if data.Index(i).Interface().(string) != "" {
				res = append(res, data.Index(i).Interface().(string))
			}

		}
	default:
		panic("invaild parameter, hope type is []string,but input is " + data.Kind().String())
	}

	return res

}
