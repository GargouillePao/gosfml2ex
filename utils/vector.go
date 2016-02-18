package utils

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"math"
)

type vectorUtil struct {
}

func (v vectorUtil) MultiplyV2(point interface{}, num float32) (sf.Vector2f, error) {
	value, err := toVector2f(point)
	if err != nil {
		return v.Vector2Zero(), err
	}

	value.X *= num
	value.Y *= num
	return value, nil
}
func (v vectorUtil) MultiplyV3(point sf.Vector3f, num float32) sf.Vector3f {
	point.X *= num
	point.Y *= num
	point.Z *= num
	return point
}
func (v vectorUtil) DistanceV1(point1 interface{}, point2 interface{}) (output float32, err error) {
	output = 0
	value1 := 0.0
	value2 := 0.0

	value1, err = toFloat64(point1)
	if err != nil {
		return
	}
	value2, err = toFloat64(point2)
	if err != nil {
		return
	}

	output = float32(math.Abs(value1 - value2))
	return
}
func (v vectorUtil) DistanceV2(point1 interface{}, point2 interface{}) (output float32, err error) {
	output = 0
	var value1 sf.Vector2f
	var value2 sf.Vector2f
	value1, err = toVector2f(point1)
	if err != nil {
		return
	}
	value2, err = toVector2f(point2)
	if err != nil {
		return
	}
	output = v.DistanceV2f(value1, value2)
	return
}
func (v vectorUtil) DistanceV2f(point1, point2 sf.Vector2f) (output float32) {
	vector := point2.Minus(point1)
	output = float32(math.Sqrt(math.Pow(float64(vector.X), 2) + math.Pow(float64(vector.Y), 2)))
	return
}
func (v vectorUtil) DistanceV3(point1 sf.Vector3f, point2 sf.Vector3f) (output float32, err error) {
	vector := sf.Vector3f{X: point1.X - point2.X, Y: point1.Y - point2.Y, Z: point1.Z - point2.Z}
	output = float32(math.Sqrt(math.Pow(float64(vector.X), 2) + math.Pow(float64(vector.X), 2) + math.Pow(float64(vector.Z), 2)))
	err = nil
	return
}
func (v *vectorUtil) LerpFloat32(value1 float32, value2 float32, lerp float32) float32 {
	return lerpV1(value1, value2, lerp)
}
func (v vectorUtil) LerpV1(point1 interface{}, point2 interface{}, lerp float32) (output float32, err error) {
	output = 0
	value1 := 0.0
	value2 := 0.0

	value1, err = toFloat64(point1)
	if err != nil {
		return
	}
	value2, err = toFloat64(point2)
	if err != nil {
		return
	}

	output = lerpV1(float32(value1), float32(value2), lerp)
	return
}
func (v vectorUtil) LerpV2(point1 interface{}, point2 interface{}, lerp float32) (output sf.Vector2f, err error) {
	var value1 sf.Vector2f
	var value2 sf.Vector2f
	value1, err = toVector2f(point1)
	if err != nil {
		return
	}
	value2, err = toVector2f(point2)
	if err != nil {
		return
	}
	output = sf.Vector2f{
		X: lerpV1(value1.X, value2.X, lerp),
		Y: lerpV1(value1.Y, value2.Y, lerp),
	}
	return
}
func (v vectorUtil) LerpV3(point1 interface{}, point2 interface{}, lerp float32) (output sf.Vector3f, err error) {
	var value1 sf.Vector3f
	var value2 sf.Vector3f
	var ok bool
	value1, ok = point1.(sf.Vector3f)
	if !ok {
		err = NewError(Errors.CannotMius)
		return
	}
	value2, ok = point2.(sf.Vector3f)
	if !ok {
		err = NewError(Errors.NotTheSameType)
		return
	}
	output = sf.Vector3f{
		X: lerpV1(value1.X, value2.X, lerp),
		Y: lerpV1(value1.Y, value2.Y, lerp),
		Z: lerpV1(value1.Z, value2.Z, lerp),
	}
	return
}
func (v vectorUtil) LerpV4(point1 interface{}, point2 interface{}, lerp float32) (output sf.Color, err error) {
	var value1 sf.Color
	var value2 sf.Color
	var ok bool
	value1, ok = point1.(sf.Color)
	if !ok {
		err = NewError(Errors.CannotMius)
		return
	}
	value2, ok = point2.(sf.Color)
	if !ok {
		err = NewError(Errors.NotTheSameType)
		return
	}
	output = sf.Color{
		R: lerpV1UInt8(value1.R, value2.R, lerp),
		G: lerpV1UInt8(value1.G, value2.G, lerp),
		B: lerpV1UInt8(value1.B, value2.B, lerp),
		A: lerpV1UInt8(value1.A, value2.A, lerp),
	}
	return
}
func lerpV1UInt8(value1 uint8, value2 uint8, lerp float32) uint8 {
	return uint8(float32(value1)*(1-lerp) + float32(value2)*lerp)
}
func lerpV1(value1 float32, value2 float32, lerp float32) float32 {
	return value1*(1-lerp) + value2*lerp
}
func (v vectorUtil) LerpV2f(point1 sf.Vector2f, point2 sf.Vector2f, lerp float32) (output sf.Vector2f) {
	output = sf.Vector2f{
		X: lerpV1(point1.X, point2.X, lerp),
		Y: lerpV1(point1.Y, point2.Y, lerp),
	}
	return
}
func isV1(value interface{}) bool {
	if _, err := toFloat64(value); err != nil {
		return false
	}
	return true
}
func isV2(value interface{}) bool {
	if _, err := toVector2f(value); err != nil {
		return false
	}
	return true
}
func (v *vectorUtil) JudageDimension(value interface{}) uint8 {
	if isV1(value) {
		return 1
	}
	if isV2(value) {
		return 2
	}
	if _, ok := value.(sf.Vector3f); ok {
		return 3
	}
	if _, ok := value.(sf.Color); ok {
		return 4
	}
	return 0
}
func toFloat64(point interface{}) (output float64, err error) {
	if point == nil {
		err = NewError(Errors.NotTheSameType)
		return
	}
	switch value := point.(type) {
	case int:
		output = float64(value)
	case int32:
		output = float64(value)
	case int64:
		output = float64(value)
	case float32:
		output = float64(value)
	case float64:
		output = value
	default:
		err = NewError(Errors.CannotMius)
	}
	return
}
func toVector2f(point interface{}) (output sf.Vector2f, err error) {
	if point == nil {
		err = NewError(Errors.NilAttribute)
		return
	}
	switch value := point.(type) {
	case sf.Vector2i:
		output = sf.Vector2f{float32(value.X), float32(value.Y)}
	case sf.Vector2f:
		output = value
	case sf.Vector2u:
		output = sf.Vector2f{float32(value.X), float32(value.Y)}
	default:
		err = NewError(Errors.CannotMius)
	}
	return
}
func (v vectorUtil) Vector2Zero() sf.Vector2f {
	return sf.Vector2f{0, 0}
}
func (v vectorUtil) Vector3Zero() sf.Vector3f {
	return sf.Vector3f{0, 0, 0}
}
func (v vectorUtil) NorV2(vector sf.Vector2f) sf.Vector2f {
	length := v.NormV2(vector)
	return sf.Vector2f{X: vector.Y / -length, Y: vector.X / length}
}
func (v vectorUtil) DirV2(vector sf.Vector2f) sf.Vector2f {
	length := v.NormV2(vector)
	return sf.Vector2f{X: vector.X / length, Y: vector.Y / length}
}
func (v vectorUtil) NormV2(vector sf.Vector2f) float32 {
	length := float32(math.Sqrt(math.Pow(float64(vector.X), 2) + math.Pow(float64(vector.Y), 2)))
	return length
}
func (v vectorUtil) AngleV2(vector1 sf.Vector2f, vector2 sf.Vector2f) float32 {
	norm1 := v.NormV2(vector1)
	norm2 := v.NormV2(vector2)
	t := vector1.X*vector2.X + vector1.Y*vector2.Y
	if norm1*norm2 != 0 {
		cos := t / (norm1 * norm2)
		return 180 * float32(math.Acos(float64(cos))) / math.Pi
	}
	return 0
}
