package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"runtime"
	"sfmlex/animations"
	shapeEX "sfmlex/shapes"
	sfUtil "sfmlex/utils"
	"time"
)

var vectorUtil = sfUtil.Utils().Vector

func init() {
	runtime.LockOSThread()
}

func main() {
	ticker := time.NewTicker(time.Second / 30)
	setting := sf.DefaultContextSettings()
	setting.AntialiasingLevel = 8
	renderWindow := sf.NewRenderWindow(sf.VideoMode{800, 600, 32}, "Events (GoSFML2)", sf.StyleDefault, setting)

	shape, _ := shapeEX.NewCurveCicleShape(10, 100, 8)
	shape.SetFillColor(sf.Color{255, 100, 55, 255})

	animation1 := animations.NewAnimation(1)
	clip10 := animations.NewSingleAnimationClip(
		shape.GetPosition(),
		sf.Vector2f{500, 200},
		func(step interface{}) {
			stepv2, _ := step.(sf.Vector2f)
			shape.SetPosition(stepv2)
		}, func() {
		})
	clip11 := animations.NewSingleAnimationClip(
		sf.Vector2f{500, 200},
		sf.Vector2f{200, 400},
		func(step interface{}) {
			stepv2, _ := step.(sf.Vector2f)
			shape.SetPosition(stepv2)
		}, func() {
		})
	clip10.SetFrameCount(120)
	clip11.SetFrameCount(120)
	animation1.AddClip(clip10)
	animation1.AddClip(clip11)

	animation2 := animations.NewAnimation(1)
	clip20 := animations.NewLoopAnimation(
		0.2, 1.8, -1, animations.Pingpong,
		func(step interface{}) {
			stepv1, _ := step.(float32)
			for i := 0; i < 8; i++ {
				if i%2 == 0 {
					shape.ExpendPointTo(uint(i), stepv1)
				}

			}
		}, func() {
		})
	clip20.SetAnimationCurve(func(num float32) float32 {
		return num * num
	})
	clip20.SetFrameCount(20)
	animation2.AddClip(clip20)

	animation1.Play()
	animation2.Play()
	for renderWindow.IsOpen() {
		select {
		case <-ticker.C:
			for event := renderWindow.PollEvent(); event != nil; event = renderWindow.PollEvent() {
				switch event.(type) {
				case sf.EventClosed:
					renderWindow.Close()
				}
			}
		}
		animation1.Animate()
		animation2.Animate()
		renderWindow.Clear(sf.ColorWhite())
		renderWindow.Draw(shape, sf.DefaultRenderStates())
		renderWindow.Display()
	}
}
