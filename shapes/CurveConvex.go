package shapes

import (
	sf "bitbucket.org/krepa098/gosfml2"
	sfutils "github.com/GargouillePao/gosfml2ex/utils"
	"math"
)

var sfUtil = sfutils.Utils()

type CurveConvexShape struct {
	sf.ConvexShape
	step          int
	ctrlPoints    []sf.Vector2f
	ctrlPointsOld []sf.Vector2f
	curves        int
}

func NewCurveConvexShape(step int) (cShape *CurveConvexShape, err error) {
	shape, err := sf.NewConvexShape()
	cShape = &CurveConvexShape{ConvexShape: *shape, step: step, ctrlPoints: make([]sf.Vector2f, 0), curves: 3}
	return
}
func NewCurveCicleShape(step int, radius float32, num uint) (cShape *CurveConvexShape, err error) {
	shape, err := sf.NewConvexShape()
	cShape = &CurveConvexShape{ConvexShape: *shape, step: step, ctrlPoints: make([]sf.Vector2f, 0), curves: 3}
	cShape.SetPointCount(num)
	var i uint
	for i = 0; i < num; i++ {
		cShape.SetPoint(i, sf.Vector2f{radius * float32(math.Cos(float64(i)*math.Pi*2/float64(num))), radius * float32(math.Sin(float64(i)*math.Pi*2/float64(num)))})
	}
	cShape.AllToCurve()
	return
}
func (c *CurveConvexShape) SetPointCount(count uint) {
	c.ConvexShape.SetPointCount(count * uint(c.step))
	c.ctrlPoints = make([]sf.Vector2f, count)
	c.ctrlPointsOld = make([]sf.Vector2f, count)
}
func (c *CurveConvexShape) SetPoint(index uint, point sf.Vector2f) {
	c.ctrlPoints[index] = point
	c.ctrlPointsOld[index] = point
	c.ConvexShape.SetPoint(index, point)
}
func (c *CurveConvexShape) GetPoint(index uint) (point sf.Vector2f) {
	return c.ctrlPoints[index]
}
func (c *CurveConvexShape) GetNearPointIndex(point sf.Vector2f, maxDistance float32) (index uint, hasPoint bool) {
	hasPoint = true
	distance, _ := sfUtil.Vector.DistanceV2(c.ctrlPoints[0].Plus(c.GetPosition()), point)
	index = uint(0)
	for i := 1; i < len(c.ctrlPoints); i++ {
		dis, _ := sfUtil.Vector.DistanceV2(c.ctrlPoints[i].Plus(c.GetPosition()), point)
		if distance > dis {
			distance = dis
			index = uint(i)
		}
	}
	if distance < maxDistance {
		hasPoint = false
	}
	return
}
func (c *CurveConvexShape) MovePoint(index uint, point sf.Vector2f) {
	c.ctrlPoints[index] = c.ctrlPoints[index].Plus(point)
	c.PointToCurve(int(index))
}
func (c *CurveConvexShape) ExpendPoint(index uint, expand float32) {
	c.ctrlPoints[index], _ = sfUtil.Vector.MultiplyV2(c.ctrlPoints[index], expand)
	c.PointToCurve(int(index))
}
func (c *CurveConvexShape) ExpendPointTo(index uint, expand float32) {
	c.ctrlPoints[index], _ = sfUtil.Vector.MultiplyV2(c.ctrlPointsOld[index], expand)
	c.PointToCurve(int(index))
}
func (c *CurveConvexShape) ResetPoint(index uint, lerp float32) {
	oldPoint := c.ctrlPointsOld[index]
	if c.ctrlPoints[index] != oldPoint {
		c.ctrlPoints[index], _ = sfUtil.Vector.LerpV2(c.ctrlPoints[index], oldPoint, lerp)
		c.PointToCurve(int(index))
	}
}
func (c *CurveConvexShape) ResetAllPoints(lerp float32) {
	for i := 0; i < len(c.ctrlPointsOld); i++ {
		c.ResetPoint(uint(i), lerp)
	}
}
func (c *CurveConvexShape) PointToCurve(index int) {
	ctrlPointLen := len(c.ctrlPoints)
	c.smoothCurve((index - 1 + ctrlPointLen) % ctrlPointLen)
	c.smoothCurve(index)
	c.smoothCurve((index + 1) % ctrlPointLen)
}
func (c *CurveConvexShape) toCurve(index int, ctrlPoints []sf.Vector2f) {
	for j := 0; j < c.step; j++ {
		t := float32(j) / float32(c.step)
		c.ConvexShape.SetPoint(uint(index*c.step+j), sfUtil.Graphics.BCurve(ctrlPoints, t, c.curves))
	}
}
func (c *CurveConvexShape) smoothCurve(index int) {
	ctrlPointLen := len(c.ctrlPoints)
	ctrlPoint := c.ctrlPoints[index]
	ctrlPointF := c.ctrlPoints[(index-1+ctrlPointLen)%ctrlPointLen]
	ctrlPointB := c.ctrlPoints[(index+1)%ctrlPointLen]
	ctrlPointSlice := []sf.Vector2f{ctrlPointF, ctrlPoint, ctrlPointB}
	c.toCurve(index, ctrlPointSlice)
}
func (c *CurveConvexShape) AllToCurve() {
	ctrlPointLen := len(c.ctrlPoints)
	for i := 0; i < ctrlPointLen; i++ {
		c.smoothCurve(i)
	}
}
