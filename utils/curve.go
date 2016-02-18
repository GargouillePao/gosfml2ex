package utils

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"math"
)

type graphicsUtil struct {
}

func (g graphicsUtil) BCurve(ctrlPoints []sf.Vector2f, t float32) (point sf.Vector2f) {
	point.X = 0
	point.Y = 0
	n := len(ctrlPoints)
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
		point.X += ctrlPoints[i].X * changeOut
		point.Y += ctrlPoints[i].Y * changeOut
	}
	return
}
func (g graphicsUtil) BezierCurve(ctrlPoints []sf.Vector2f, t float32) (point sf.Vector2f) {
	n := len(ctrlPoints)
	switch n {
	case 0:
	case 3:
		a := (1 - t) * (1 - t)
		b := (1 - t) * t * 2
		c := t * t
		point = sf.Vector2f{
			a*ctrlPoints[0].X + b*ctrlPoints[1].X + c*ctrlPoints[2].X,
			a*ctrlPoints[0].Y + b*ctrlPoints[1].Y + c*ctrlPoints[2].Y,
		}
	case 4:
		a := (1 - t) * (1 - t) * (1 - t)
		b := (1 - t) * (1 - t) * t * 3
		c := t * t * (1 - t) * 3
		d := t * t * t
		point = sf.Vector2f{
			a*ctrlPoints[0].X + b*ctrlPoints[1].X + c*ctrlPoints[2].X + d*ctrlPoints[3].X,
			a*ctrlPoints[0].Y + b*ctrlPoints[1].Y + c*ctrlPoints[2].Y + d*ctrlPoints[3].Y,
		}
	default:
		point = ctrlPoints[0]
	}
	return
}
