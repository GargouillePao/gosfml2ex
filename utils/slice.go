package utils

import (
	sf "bitbucket.org/krepa098/gosfml2"
)

type sliceUtil struct {
}

func (s sliceUtil) Equilongf(slice1 *[]float32, slice2 *[]float32) {
	len1 := len(*slice1)
	len2 := len(*slice2)
	if len1 == len2 {
		return
	}
	if len1 < len2 {
		for {
			if len(*slice1) == len2 {
				return
			}
			*slice1 = append(*slice1, 0)
		}
	} else {
		for {
			if len(*slice2) == len1 {
				return
			}
			*slice2 = append(*slice2, 0)
		}
	}
}
func (s sliceUtil) Lerpf(slice1 []float32, slice2 []float32, lerp float32) []float32 {
	length := len(slice1)
	if length != len(slice2) {
		return nil
	}
	slice := make([]float32, length)
	for i := 0; i < length; i++ {
		slice[i] = lerpV1(slice1[i], slice2[i], lerp)
	}
	return slice
}
func (s sliceUtil) ToVector1(slice []float32) float32 {
	if len(slice) > 0 {
		return slice[0]
	}
	return 0
}
func (s sliceUtil) ToVector2(slice []float32) sf.Vector2f {
	switch len(slice) {
	case 0:
		return sf.Vector2f{0, 0}
	case 1:
		return sf.Vector2f{slice[0], 0}
	default:
		return sf.Vector2f{slice[0], slice[1]}
	}
}
func (s sliceUtil) ToVector3(slice []float32) sf.Vector3f {
	switch len(slice) {
	case 0:
		return sf.Vector3f{0, 0, 0}
	case 1:
		return sf.Vector3f{slice[0], 0, 0}
	case 2:
		return sf.Vector3f{slice[0], slice[1], 0}
	default:
		return sf.Vector3f{slice[0], slice[1], slice[2]}
	}
}
func (s sliceUtil) ToVector4(slice []float32) sf.Color {
	switch len(slice) {
	case 0:
		return sf.Color{0, 0, 0, 0}
	case 1:
		return sf.Color{uint8(slice[0]), 0, 0, 0}
	case 2:
		return sf.Color{uint8(slice[0]), uint8(slice[1]), 0, 0}
	case 3:
		return sf.Color{uint8(slice[0]), uint8(slice[1]), uint8(slice[2]), 0}
	default:
		return sf.Color{uint8(slice[0]), uint8(slice[1]), uint8(slice[2]), uint8(slice[3])}
	}
}
