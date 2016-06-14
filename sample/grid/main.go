package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	"github.com/GargouillePao/gosfml2ex/images"
	"github.com/GargouillePao/gosfml2ex/shapes"
	sfUtil "github.com/GargouillePao/gosfml2ex/utils"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
}

var mathUtil = sfUtil.Utils.Math

func main() {
	ticker := time.NewTicker(time.Second / 30)
	setting := sf.DefaultContextSettings()
	setting.AntialiasingLevel = 8
	renderWindow := sf.NewRenderWindow(sf.VideoMode{800, 600, 32}, "Events (GoSFML2)", sf.StyleDefault, setting)

	grid, err := shapes.NewGirdShape(sf.Vector2f{500, 400}, sf.Vector2f{50, 40}, sf.ColorWhite())

	imageGot, _, err := images.ReadImage("./011.jpg", 0, 0, 500, 400)
	texture, err := images.ReadTextureFromImage(*imageGot)
	if err != nil {
		fmt.Println(err)
		return
	}
	texture.SetSmooth(true)
	grid.SetTexture(texture)
	grid.Move(sf.Vector2f{100, 100})
	grid.Rotate(15)
	var pointX, pointY int
	var gotPoint bool

	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			mousePosi := sf.MouseGetPosition(renderWindow)
			mousePosf := sf.Vector2f{float32(mousePosi.X), float32(mousePosi.Y)}
			for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
				switch event.(type) {
				case sf.EventClosed:
					renderWindow.Close()
				case sf.EventMouseButtonPressed:
					pointX, pointY, gotPoint = grid.GetNearestPointIndex(mousePosf)
				}
			}
			if sf.IsMouseButtonPressed(0) {
				if gotPoint {
					grid.SmoothPointsTo(pointX, pointY, mousePosf, 10, func(data float32) float32 {
						return (mathUtil.Sin(data, -0.5, 1)/2 + 0.5) / 5
					})
				}
			}
		}
		renderWindow.Clear(sf.ColorWhite())
		renderWindow.Draw(grid, sf.DefaultRenderStates())
		renderWindow.Display()
	}
}
