package utils

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"math"
)

type graphicsUtil struct {
}

func (g graphicsUtil) BCurve(points []sf.Vector2f, t float32, n int) (point sf.Vector2f) {
	point.X = 0
	point.Y = 0
	for i := 0; i < n; i++ {
		var changeOut float32
		for j := 0; j < n-i; j++ {
			factorialJ := mathUtil{}.Factorial(j)
			factorialN := mathUtil{}.Factorial(n - j)
			changeIn := float32(math.Pow(float64(t+float32(n-1-i-j)), float64(n-1))) * float32(n) / float32(factorialJ*factorialN)
			if j%2 != 0 {
				changeOut -= changeIn
			} else {
				changeOut += changeIn
			}
		}
		point.X += points[i].X * changeOut
		point.Y += points[i].Y * changeOut
	}
	return
}
