package shapes

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"math"
)

type StickShap struct {
	sf.ConvexShape
	ctrlPoint    []sf.Vector2f
	height       float32
	angleTan     float32
	angleCos     float32
	angleSin     float32
	width        float32
	step         float32
	radiusTop    float32
	radiusBottom float32
}

func NewStickShape(height float32, width float32, step float32) (shape *StickShap, err error) {
	convex, err := sf.NewConvexShape()
	shape = &StickShap{ConvexShape: *convex, height: height, width: width, step: step, ctrlPoint: make([]sf.Vector2f, 12), radiusTop: 0, radiusBottom: 0}
	shape.SetPointCount(uint(4*step + 4))
	shape.updateShape()
	return
}

func (s *StickShap) SetHeight(height float32, lerp float32) {
	s.height = Vector.LerpFloat32(s.height, height, lerp)
	s.updateShape()
}
func (s *StickShap) SetWidth(width float32, lerp float32) {
	s.width = Vector.LerpFloat32(s.width, width, lerp)
	s.updateShape()
}
func (s *StickShap) SetRadius(location int8, radius float32) {
	if location == 0 {
		s.radiusBottom = radius
	} else {
		s.radiusTop = radius
	}
	s.updateShape()
}
func (s *StickShap) SetAngle(angle float32) {
	s.angleTan = float32(math.Tan(math.Pi * float64(angle) / 180))
	s.angleCos = float32(math.Cos(math.Pi * float64(angle) / 180))
	s.angleSin = float32(math.Sin(math.Pi * float64(angle) / 180))
	s.updateShape()
}
func (s *StickShap) updateShape() {
	offset := s.width
	offsetWidth := s.height * s.angleTan
	waistLength := offsetWidth / s.angleSin
	offsetHeight := s.height - (waistLength-offsetWidth)*s.angleCos
	if s.width > s.height {
		offset = s.height
	}
	offsetBottom := offset / 2 * s.radiusBottom
	offsetTop := offset / 2 * s.radiusTop

	s.ctrlPoint[0] = sf.Vector2f{-s.width/2 + offsetBottom, 0}
	s.ctrlPoint[1] = sf.Vector2f{-s.width / 2, 0}
	s.ctrlPoint[2] = sf.Vector2f{-s.width / 2, -offsetBottom}

	s.ctrlPoint[3] = sf.Vector2f{-s.width/2 - (s.height-offsetHeight-offsetTop)*s.angleTan, -s.height + offsetHeight + offsetTop}
	s.ctrlPoint[4] = sf.Vector2f{-s.width/2 - offsetWidth, -s.height}
	s.ctrlPoint[5] = sf.Vector2f{-s.width/2 + offsetTop, -s.height}

	s.ctrlPoint[6] = sf.Vector2f{s.width/2 - offsetTop, -s.height}
	s.ctrlPoint[7] = sf.Vector2f{s.width/2 + offsetWidth, -s.height}
	s.ctrlPoint[8] = sf.Vector2f{s.width/2 + (s.height-offsetHeight-offsetTop)*s.angleTan, -s.height + offsetHeight + offsetTop}

	s.ctrlPoint[9] = sf.Vector2f{s.width / 2, -offsetBottom}
	s.ctrlPoint[10] = sf.Vector2f{s.width / 2, 0}
	s.ctrlPoint[11] = sf.Vector2f{s.width/2 - offsetBottom, 0}

	s.toCurve(0, s.ctrlPoint[0:3])
	s.toCurve(1, s.ctrlPoint[3:6])
	s.toCurve(2, s.ctrlPoint[6:9])
	s.toCurve(3, s.ctrlPoint[9:12])
}
func (c *StickShap) toCurve(index int, ctrlPoints []sf.Vector2f) {
	for j := 0; j <= int(c.step); j++ {
		t := float32(j) / float32(c.step)
		pointIndex := uint(index*int(c.step) + j)
		if pointIndex < c.GetPointCount() {
			c.SetPoint(pointIndex, sfUtil.Graphics.BezierCurve(ctrlPoints, t))
		}
	}
}
