package utils

type mathUtil struct {
}

func (m mathUtil) Factorial(num int) int {
	_, out := factorial(num, num)
	return out
}

func factorial(index int, num int) (int, int) {

	index--
	num *= index
	if index <= 0 {
		return 0, 1
	}
	if index <= 1 {
		return 0, num
	}
	return factorial(index, num)
}
