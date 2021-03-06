package animations

import (
	sf "bitbucket.org/krepa098/gosfml2"
	shapeEX "github.com/GargouillePao/gosfml2ex/shapes"
)

type PathTracker struct {
	step      float32
	offset    float32
	stepCount float32
	rotateNor sf.Vector2f
	transform sf.Transformer
	path      shapeEX.PathVertexts
}

func NewPathTracker(transform sf.Transformer, path shapeEX.PathVertexts) *PathTracker {
	maxStep := float32(path.GetVertexCount())
	return &PathTracker{step: 0, offset: 0, stepCount: maxStep, transform: transform, path: path}
}
func (p *PathTracker) SetRoateNor(x float32, y float32) {
	p.rotateNor = sf.Vector2f{x, y}
}
func (p *PathTracker) SetStep(step float32) {
	p.step = step
	pointIndex := int(p.step * p.stepCount)
	if step > 0 && step < 1 {

		moveStep := p.step*p.stepCount - float32(pointIndex)
		point, err := p.path.GetVertext(pointIndex)

		pointNext, err := p.path.GetVertext(pointIndex + 1)
		if err != nil {
			return
		}
		pos := Vector.LerpV2f(point.Position, pointNext.Position, moveStep)
		p.transform.SetPosition(pos)
	}

	nor, err := p.path.GetVertextNormal(pointIndex)
	if err != nil {
		return
	}
	angle := Vector.AngleV2(nor, p.rotateNor)
	p.transform.SetRotation(angle)
}
