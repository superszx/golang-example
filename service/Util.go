package service1

type Util struct {
}

func (Util) Swap(num1, num2 *int) {
	*num1, *num2 = *num2, *num1
}
