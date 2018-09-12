package Linq

type Enumable struct {
	source interface{}
}

func New(source interface{}) Enumable {
	return Enumable{source}
}
