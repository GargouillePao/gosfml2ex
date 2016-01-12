package shapes

import (
	sf "bitbucket.org/krepa098/gosfml2"
	sfutils "github.com/GargouillePao/gosfml2ex/utils"
)

var Errors = sfutils.Errors()
var Vector = sfutils.Utils().Vector
var Grapics = sfutils.Utils().Graphics

type PathVertexts interface {
	GetVertextNormal(index int) (sf.Vector2f, error)
	GetVertext(index int) (sf.Vertex, error)
	GetVertextDirection(index int) (sf.Vector2f, error)
	GetVertexCount() int
}

type PathShape struct {
	sf.VertexArray
	ctrls []Lever
	step  float32
}

func NewPathShap(step float32) *PathShape {
	pathShap := &PathShape{step: step}
	pathShap.PrimitiveType = sf.PrimitiveLinesStrip
	return pathShap
}

func (p *PathShape) SetCtrl(index int, vertex sf.Vertex) {
	var radius float32
	for len(p.ctrls) <= index {
		lastOne, err := p.GetCtrl(index - 1)
		if err != nil {
			radius = 25
		} else {
			radius = Vector.NormV2(vertex.Position.Minus(lastOne.GetPoint(0))) / 2
		}
		p.ctrls = append(p.ctrls, GetLever(vertex.Position, radius))
		for i := 0; i < int(p.step); i++ {
			p.Append(vertex)
		}
	}
}
func (p *PathShape) SetEndCtrl(index int, vertex sf.Vertex) {
	var radius float32
	for len(p.ctrls) <= index {
		lastOne, err := p.GetCtrl(index - 1)
		if err != nil {
			radius = 25
		} else {
			radius = Vector.NormV2(vertex.Position.Minus(lastOne.GetPoint(0))) / 2
		}
		p.ctrls = append(p.ctrls, GetLever(vertex.Position, radius))
		p.Append(vertex)
	}
}

func (p *PathShape) AllToCurve() {
	for i := 0; i < len(p.ctrls); i++ {
		p.toCurve(i)
	}
}

func (p *PathShape) toCurve(index int) {
	ctrlLen := len(p.ctrls)
	index1 := (index - 1 + ctrlLen) % ctrlLen
	index2 := index
	index3 := (index + 1) % ctrlLen
	index4 := (index + 2) % ctrlLen
	lever1, err := p.GetCtrl(index1)
	if err != nil {
		return
	}
	lever2, err := p.GetCtrl(index2)
	if err != nil {
		return
	}
	lever3, err := p.GetCtrl(index3)
	if err != nil {
		return
	}
	lever4, err := p.GetCtrl(index4)
	if err != nil {
		return
	}
	lever2.SetDirection(lever1.GetPoint(0), lever3.GetPoint(0))
	lever3.SetDirection(lever2.GetPoint(0), lever4.GetPoint(0))
	p.ctrls[index2] = lever2
	p.ctrls[index3] = lever3
	ctrlPoints := []sf.Vector2f{
		p.ctrls[index2].GetPoint(0),
		p.ctrls[index2].GetPoint(1),
		p.ctrls[index3].GetPoint(-1),
		p.ctrls[index3].GetPoint(0),
	}
	for i := 0; i < int(p.step); i++ {
		t := float32(i) / p.step
		vertexIndex := i + index2*int(p.step)
		vertex, _ := p.GetVertext(vertexIndex)
		vertex.Position = Grapics.BezierCurve(ctrlPoints, t)
		p.SetVertext(vertexIndex, vertex)
	}
}

func (p *PathShape) GetVertext(index int) (sf.Vertex, error) {
	if p.GetVertexCount() > index && index >= 0 {
		return p.Vertices[index], nil
	}
	return sf.Vertex{}, sfutils.NewError(Errors.OutOfRange)
}
func (p *PathShape) SetVertext(index int, vertex sf.Vertex) {
	if index < p.GetVertexCount() {
		p.Vertices[index] = vertex
	}
}
func (p *PathShape) GetCtrl(index int) (Lever, error) {
	if index < len(p.ctrls) && index >= 0 {
		return p.ctrls[index], nil
	}
	return Lever{}, sfutils.NewError(Errors.OutOfRange)
}

func (p *PathShape) GetVertextNormal(index int) (sf.Vector2f, error) {
	v, err := p.GetVertextDirection(index)
	if err != nil {
		return sf.Vector2f{}, err
	}
	return Vector.NorV2(v), nil
}

func (p *PathShape) GetVertextDirection(index int) (sf.Vector2f, error) {
	vertex1, err := p.GetVertext((index - 1 + p.GetVertexCount()) % p.GetVertexCount())
	if err != nil {
		return Vector.Vector2Zero(), err
	}
	vertex2, err := p.GetVertext((index + 1) % p.GetVertexCount())
	if err != nil {
		return Vector.Vector2Zero(), err
	}
	return vertex1.Position.Minus(vertex2.Position), nil
}

type Lever struct {
	point      sf.Vector2f
	ctrlPoint1 sf.Vector2f
	ctrlPoint2 sf.Vector2f
	radius1    float32
	radius2    float32
}

func NewLever(point sf.Vector2f, radius float32) *Lever {
	return &Lever{point: point, radius1: radius, radius2: radius}
}
func GetLever(point sf.Vector2f, radius float32) Lever {
	return Lever{point: point, radius1: radius, radius2: radius}
}

func (l *Lever) SetDirection(point1 sf.Vector2f, point2 sf.Vector2f) {
	direction := point2.Minus(point1)
	direction = Vector.DirV2(direction)
	dir1 := direction
	dir2 := direction
	dir1.X *= l.radius1
	dir1.Y *= l.radius1
	dir2.X *= l.radius2
	dir2.Y *= l.radius2
	l.ctrlPoint1 = l.point.Minus(dir1)
	l.ctrlPoint2 = l.point.Plus(dir2)
}

func (l *Lever) GetPoint(num int8) sf.Vector2f {
	switch {
	case num == 0:
		return l.point
	case num > 0:
		return l.ctrlPoint2
	default:
		return l.ctrlPoint1
	}
}

func (l *Lever) SetRadius(num int8, radius float32) {
	if num > 0 {
		l.radius2 = radius
	} else {
		l.radius1 = radius
	}
}
