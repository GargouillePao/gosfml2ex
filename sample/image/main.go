package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"fmt"
	"github.com/GargouillePao/gosfml2ex/animations"
	"github.com/GargouillePao/gosfml2ex/images"
	"math"
	"runtime"
	"time"
)

func init() {
	runtime.LockOSThread()
}

func main() {

	ticker := time.NewTicker(time.Second / 30)
	setting := sf.DefaultContextSettings()
	setting.AntialiasingLevel = 8
	renderWindow := sf.NewRenderWindow(sf.VideoMode{800, 600, 32}, "Events (GoSFML2)", sf.StyleDefault, setting)

	round, err := sf.NewCircleShape()

	imageGot, _, err := images.ReadImage("./011.jpg", -50, -50, 600, 600)
	if err != nil {
		fmt.Println(err)
		return
	}
	imageGot__ := images.ImageEnhanceRGBWithFunc(imageGot, func(r, g, b float32) (outr, outg, outb float32) {
		outr = 127 * float32(math.Sin(float64(r-127)/255*math.Pi+1))
		outg = g
		outb = b
		return
	})
	texture, err := images.ReadTextureFromImage(*imageGot__)
	if err != nil {
		fmt.Println(err)
		return
	}
	texture.SetSmooth(true)
	round.SetPosition(sf.Vector2f{100, 50})
	round.SetRadius(300)
	round.SetTexture(texture, true)

	clip := animations.NewLoopAnimation([]float32{0}, []float32{255}, -1, animations.Pingpong, func(values []float32) {
		imageGot__ = images.ImageEnhanceRGBWithFunc(imageGot,
			func(r, g, b float32) (outr, outg, outb float32) {
				outr = 127 * float32(math.Sin(float64(r-127+values[0])/255*math.Pi+1))
				outg = g
				outb = b
				return
			})
		texture, _ := images.ReadTextureFromImage(*imageGot__)
		round.SetTexture(texture, true)
	}, nil)
	clip.SetFrameCount(60)
	animation := animations.NewAnimation(0)
	animation.AddClip(clip)
	animation.Play()
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
		animation.Animate()
		renderWindow.Clear(sf.ColorWhite())
		renderWindow.Draw(round, sf.DefaultRenderStates())
		renderWindow.Display()
	}
}
