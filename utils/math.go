package utils

import (
	"math"
)

type mathUtil struct {
}

func (m mathUtil) Factorial(num int) int {
	_, out := factorial(num, num)
	return out
}

func (m mathUtil) Power(num float32, index float32) float32 {
	return float32(math.Pow(float64(num), float64(index)))
}
func (m mathUtil) Sqrt(num float32) float32 {
	var xhalf float32 = 0.5 * num // get bits for floating VALUE
	i := math.Float32bits(num)    // gives initial guess y0
	i = 0x5f375a86 - (i >> 1)     // convert bits BACK to float
	num = math.Float32frombits(i) // Newton step, repeating increases accuracy
	num = num * (1.5 - xhalf*num*num)
	num = num * (1.5 - xhalf*num*num)
	num = num * (1.5 - xhalf*num*num)
	return 1 / num
}
func (m mathUtil) QuardraticY(num float32, centerX float32, centerY float32, radiusX float32) float32 {
	b := centerY
	t := centerX
	a := -1 * b / ((radiusX - t) * (radiusX - t))
	return a*(num-t)*(num-t) + b
}
func (m mathUtil) SymmetricY(num float32, slope float32, centerX float32) float32 {
	return float32(math.Abs(float64((num - centerX) * slope)))
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
