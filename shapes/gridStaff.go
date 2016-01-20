package shapes

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	"math"
)

type GridShape struct {
	sf.VertexArray
	size         sf.Vector2f
	count        sf.Vector2f
	cellSize     sf.Vector2f
	renderStates sf.RenderStates
	pointsOld    [][]sf.Vector2f
	center       sf.Vector2f
}

func NewGirdShape(size, count sf.Vector2f, color sf.Color) (shape *GridShape, err error) {
	varray, err := sf.NewVertexArray()
	shape = &GridShape{VertexArray: *varray, size: size, count: count, renderStates: sf.DefaultRenderStates()}
	shape.PrimitiveType = sf.PrimitiveQuads
	shape.cellSize = sf.Vector2f{size.X / count.X, size.Y / count.Y}
	shape.Vertices = make([]sf.Vertex, int(count.X)*int(count.Y)*4)

	shape.pointsOld = make([][]sf.Vector2f, int(count.Y))
	for i := 0; i < int(count.Y); i++ {
		shape.pointsOld[i] = make([]sf.Vector2f, int(count.X))
	}

	shape.ForEach(1, 1, func(x, y int) {
		vertex := sf.Vertex{}
		vertex.Position = sf.Vector2f{X: float32(x) * shape.cellSize.X, Y: float32(y) * shape.cellSize.Y}
		vertex.TexCoords = sf.Vector2f{X: float32(x) * shape.cellSize.X, Y: float32(y) * shape.cellSize.Y}
		vertex.Color = color
		shape.SetVertexs(x, y, vertex)
		if x < int(shape.count.X) && y < int(shape.count.Y) {
			shape.pointsOld[y][x] = vertex.Position
		}
	})
	shape.center = sf.Vector2f{size.X / 2, size.Y / 2}
	return
}

func (g *GridShape) ForEach(xOffset, yOffset int, callback func(x, y int)) {
	for y := 0; y < int(g.count.Y)+yOffset; y++ {
		for x := 0; x < int(g.count.X)+xOffset; x++ {
			callback(x, y)
		}
	}
}

func (g *GridShape) GetNearestPointIndex(point sf.Vector2f) (xTargIndex, yTargIndex int, got bool) {
	var distanceMin float32 = 10000
	got = false
	g.ForEach(0, 0, func(x, y int) {
		index := (x + y*int(g.count.X)) * 4
		vertex := g.Vertices[index]
		verPos := g.renderStates.Transform.TransformPoint(vertex.Position)
		distance := Vector.DistanceV2f(point, verPos)
		if distance < distanceMin {
			distanceMin = distance
			xTargIndex = x
			yTargIndex = y
			got = true
			fmt.Println("DIS", point, verPos, distanceMin, Vector.DistanceV2f(point, verPos))
		}
	})
	fmt.Println("Selected", xTargIndex, yTargIndex)
	return
}

func (g *GridShape) MovePointTo(x, y int, to sf.Vector2f, lerp float32) {
	vertex, got := g.getCuadVertex(x, y, 0)
	if got {
		oldPos := g.pointsOld[y][x]
		vertex.Position = Vector.LerpV2f(oldPos, to, lerp)
		g.SetVertexs(x, y, vertex)
	}
}

func (g *GridShape) SmoothPointsTo(centerX, centerY int, to sf.Vector2f, ranges int, smoothFunction func(float32) float32) {
	xlow := centerX - ranges
	xhigh := centerX + ranges
	ylow := centerY - ranges
	yhigh := centerY + ranges
	for y := ylow; y <= yhigh; y++ {
		for x := xlow; x <= xhigh; x++ {
			distance := float32(math.Abs(float64(x-centerX))+math.Abs(float64(y-centerY))) / float32(ranges+ranges)
			g.MovePointTo(x, y, to, smoothFunction(1-distance))
		}
	}
}

func (g *GridShape) SetVertexs(xIndex, yIndex int, vertex sf.Vertex) {
	g.setCuadVertex(xIndex-1, yIndex-1, 2, vertex)
	g.setCuadVertex(xIndex, yIndex-1, 3, vertex)
	g.setCuadVertex(xIndex-1, yIndex, 1, vertex)
	g.setCuadVertex(xIndex, yIndex, 0, vertex)
}
func (g *GridShape) GetVertexs(xIndex, yIndex int) []sf.Vertex {
	ltCuadIndex := (xIndex-1+(yIndex-1)*int(g.count.X))*4 + 2
	rtCuadIndex := (xIndex+(yIndex-1)*int(g.count.X))*4 + 3
	lbCuadIndex := (xIndex-1+yIndex*int(g.count.X))*4 + 1
	rbCuadIndex := (xIndex + yIndex*int(g.count.X)) * 4
	vertexs := make([]sf.Vertex, 0)
	vertex1, got := g.getVertex(ltCuadIndex)
	if got {
		vertexs = append(vertexs, vertex1)
	}
	vertex2, got := g.getVertex(rtCuadIndex)
	if got {
		vertexs = append(vertexs, vertex2)
	}
	vertex3, got := g.getVertex(lbCuadIndex)
	if got {
		vertexs = append(vertexs, vertex3)
	}
	vertex4, got := g.getVertex(rbCuadIndex)
	if got {
		vertexs = append(vertexs, vertex4)
	}
	return vertexs
}

func (g *GridShape) getCuadVertex(xIndex, yIndex, offset int) (vertex sf.Vertex, got bool) {
	got = false
	if xIndex >= 0 && xIndex < int(g.count.X) {
		if yIndex >= 0 && yIndex < int(g.count.Y) {
			index := (xIndex + yIndex*int(g.count.X)) * 4
			vertex = g.Vertices[index+offset]
			got = true
		}
	}
	return
}

func (g *GridShape) setCuadVertex(xIndex, yIndex, offset int, vertex sf.Vertex) {
	if xIndex >= 0 && xIndex < int(g.count.X) {
		if yIndex >= 0 && yIndex < int(g.count.Y) {
			index := (xIndex + yIndex*int(g.count.X)) * 4
			g.Vertices[index+offset] = vertex
		}
	}
}

func (g *GridShape) Move(point sf.Vector2f) {
	g.renderStates.Transform.Translate(point.X, point.Y)
	g.center = g.renderStates.Transform.TransformPoint(g.center)
}
func (g *GridShape) Rotate(angle float32) {
	g.renderStates.Transform.RotateWithCenter(angle, g.center.X, g.center.Y)
}
func (g *GridShape) Scale(scaleX float32, scaleY float32) {
	g.renderStates.Transform.ScaleWithCenter(scaleX, scaleY, g.center.X, g.center.Y)
}

func (g *GridShape) setVertex(index int, vertex sf.Vertex) {
	if index >= 0 && index < len(g.Vertices) {
		g.Vertices[index] = vertex
	}
}

func (g *GridShape) getVertex(index int) (vertex sf.Vertex, got bool) {
	if index >= 0 && index < len(g.Vertices) {
		vertex = g.Vertices[index]
		got = true
	} else {
		got = false
	}
	return
}
func (g *GridShape) SetTexture(texture *sf.Texture) {
	g.renderStates.Texture = texture
}
func (g *GridShape) SetTransform(transform sf.Transform) {
	g.renderStates.Transform = transform
}
func (g *GridShape) Draw(target sf.RenderTarget, renderStates sf.RenderStates) {
	g.VertexArray.Draw(target, g.renderStates)
}
