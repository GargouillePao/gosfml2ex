package images

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"math"
	"sync"
	"time"
)

func GetImageGray(img *image.RGBA, offset image.Point) *image.Gray {
	imgGray := image.NewGray(img.Rect)
	draw.Draw(imgGray, imgGray.Rect, img.SubImage(img.Rect), offset, draw.Src)
	return imgGray
}
func GetImageGrey(img *image.RGBA, offset image.Point) *image.RGBA {
	imgGray := GetImageGray(img, offset)
	imgGrey := image.NewRGBA(img.Rect)
	draw.Draw(imgGrey, imgGrey.Rect, imgGray.SubImage(imgGray.Rect), offset, draw.Src)
	return imgGrey
}
func FilterAve(count int) [][]float32 {
	filter := [][]float32{}
	for i := 0; i < count; i++ {
		child := []float32{}
		for j := 0; j < count; j++ {
			child = append(child, 1/float32(count*count))
		}
		filter = append(filter, child)
	}
	return filter
}
func FilterSob(strength float32, isVertical bool) [][]float32 {
	if isVertical {
		return [][]float32{
			{strength, strength * float32(math.Sqrt2), strength},
			{0, 0, 0},
			{-strength, -strength * float32(math.Sqrt2), -strength},
		}
	} else {
		return [][]float32{
			{strength, 0, -strength},
			{strength * float32(math.Sqrt2), 0, -strength * float32(math.Sqrt2)},
			{strength, 0, -strength},
		}
	}
}
func ImageFilterRGB(img *image.RGBA, filter [][]float32, lowQuality bool) *image.RGBA {
	if lowQuality {
		imgAva := ImageEnhanceRGB(img, filter[0][0], 0, [3]bool{true, true, true})
		return imfilter(filter, img.Rect, func(xpos, ypos int) (r0, g0, b0, a0 uint8) {
			rgba := imgAva.RGBAAt(xpos, ypos)
			return rgba.R, rgba.G, rgba.B, rgba.A
		}, func(filterValue, colorValue float32) float32 {
			return colorValue
		})
	}
	return imfilter(filter, img.Rect, func(xpos, ypos int) (r0, g0, b0, a0 uint8) {
		rgba := img.RGBAAt(xpos, ypos)
		return rgba.R, rgba.G, rgba.B, rgba.A
	}, func(filterValue, colorValue float32) float32 {
		return filterValue * colorValue
	})
}
func ImageEnhanceRGB(img *image.RGBA, enhanceDegress float32, enhanceOffset float32, enhanceChannel [3]bool) *image.RGBA {
	rgbFunc := func(xpos, ypos int) (r0, g0, b0, a0 uint8) {
		rgba := img.RGBAAt(xpos, ypos)
		return rgba.R, rgba.G, rgba.B, rgba.A
	}
	imgEnhanced := imenhance(enhanceDegress, enhanceOffset, enhanceChannel, false, img.Rect, rgbFunc)
	return imgEnhanced
}
func ImageEnhanceRGBWithFunc(img *image.RGBA, enhanceFunc func(r, g, b float32) (float32, float32, float32)) *image.RGBA {
	rgbFunc := func(xpos, ypos int) (r0, g0, b0, a0 uint8) {
		rgba := img.RGBAAt(xpos, ypos)
		r, g, b := enhanceFunc(float32(rgba.R), float32(rgba.G), float32(rgba.B))
		return oneColorCorrect(r), oneColorCorrect(g), oneColorCorrect(b), rgba.A
	}
	imgEnhanced := imenhance(1, 0, [3]bool{true, true, true}, true, img.Rect, rgbFunc)
	return imgEnhanced
}
func ImageFilterGrey(img *image.Gray, filter [][]float32, lowQuality bool) *image.RGBA {
	if lowQuality {
		imgAva := ImageEnhanceGrey(img, filter[0][0], 0)
		return imfilter(filter, img.Rect, func(xpos, ypos int) (r0, g0, b0, a0 uint8) {
			gray := imgAva.RGBAAt(xpos, ypos)
			return gray.R, gray.R, gray.R, gray.A
		}, func(filterValue, colorValue float32) float32 {
			return colorValue
		})
	}
	return imfilter(filter, img.Rect, func(xpos, ypos int) (r0, g0, b0, a0 uint8) {
		gray := img.GrayAt(xpos, ypos)
		return gray.Y, gray.Y, gray.Y, gray.Y
	}, func(filterValue, colorValue float32) float32 {
		return filterValue * colorValue
	})
}
func ImageEnhanceGrey(img *image.Gray, enhanceDegress float32, enhanceOffset float32) *image.RGBA {
	rgbFunc := func(xpos, ypos int) (r0, g0, b0, a0 uint8) {
		gray := img.GrayAt(xpos, ypos)
		return gray.Y, gray.Y, gray.Y, gray.Y
	}
	imgEnhanced := imenhance(enhanceDegress, enhanceOffset, [3]bool{true, true, true}, false, img.Rect, rgbFunc)
	return imgEnhanced
}
func imenhance(enhanceDegress float32, enhanceOffset float32, enhanceChannel [3]bool, withFunc bool, rect image.Rectangle, rgbFunc func(xpos, ypos int) (r0, g0, b0, a0 uint8)) *image.RGBA {
	timeOld := time.Now()
	wg := sync.WaitGroup{}
	newImage := image.NewRGBA(rect)
	size := rect.Size()
	fX := size.X
	fY := size.Y

	threadCount := 8
	sizeoffset := fX / threadCount
	wg.Add(threadCount)
	for i := 0; i < threadCount; i++ {
		go func(i int) {
			for x := sizeoffset * i; x < sizeoffset*(i+1); x++ {
				for y := 0; y < fY; y++ {
					r, g, b, a := rgbFunc(x, y)
					if !withFunc {
						if enhanceChannel[0] {
							r = oneColorCorrect(float32(r)*enhanceDegress + enhanceOffset)
						}
						if enhanceChannel[1] {
							g = oneColorCorrect(float32(g)*enhanceDegress + enhanceOffset)
						}
						if enhanceChannel[2] {
							b = oneColorCorrect(float32(b)*enhanceDegress + enhanceOffset)
						}
					}
					newColor := color.RGBA{R: r, G: g, B: b, A: a}
					newImage.SetRGBA(x, y, newColor)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Enhance Time:", time.Now().Sub(timeOld))
	return newImage
}
func oneColorCorrect(r float32) (outr uint8) {
	switch {
	case r > 255:
		outr = 255
	case r < 0:
		outr = 0
	default:
		outr = uint8(r)
	}
	return
}
func imfilter(filter [][]float32, rect image.Rectangle, rgbFunc func(xpos, ypos int) (r0, g0, b0, a0 uint8), mergFunc func(a, b float32) float32) *image.RGBA {
	wg := sync.WaitGroup{}
	size := rect.Size()
	imgGrey := image.NewRGBA(rect)
	fX := len(filter)
	fY := len(filter[0])
	xoffset := fX / 2
	yoffset := fY / 2
	threadCount := (fX + fY) / 2
	sizeoffset := (size.X - fX) / threadCount

	timeStart := time.Now()

	wg.Add(threadCount)
	for i := 0; i < threadCount; i++ {
		go func(i int) {
			for x := xoffset + i*sizeoffset; x < xoffset+(i+1)*sizeoffset; x++ {
				for y := yoffset; y < size.Y-yoffset; y++ {
					newColor := imfilterMerg(filter, fX, fY, x, y, xoffset, yoffset, rgbFunc, mergFunc)
					imgGrey.SetRGBA(x, y, newColor)
				}
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("Filt Time", time.Now().Sub(timeStart))
	return imgGrey
}
func imfilterMerg(filter [][]float32, fX, fY, x, y, xoffset, yoffset int, rgbFunc func(xpos, ypos int) (r0, g0, b0, a0 uint8), mergFunc func(a, b float32) float32) color.RGBA {
	_, _, _, a := rgbFunc(x, y)
	newColor := color.RGBA{R: 0, G: 0, B: 0, A: a}
	for xx := 0; xx < fX; xx++ {
		for yy := 0; yy < fY; yy++ {
			r, g, b, _ := rgbFunc(x+xx-xoffset, y+yy-yoffset)
			newColor.R = oneColorCorrect(float32(newColor.R) + mergFunc(filter[xx][yy], float32(r)))
			if r != g {
				newColor.G = oneColorCorrect(float32(newColor.G) + mergFunc(filter[xx][yy], float32(g)))
			} else {
				newColor.G = newColor.R
			}
			switch b {
			case r:
				newColor.B = newColor.R
			case g:
				newColor.B = newColor.G
			default:
				newColor.B = oneColorCorrect(float32(newColor.B) + mergFunc(filter[xx][yy], float32(b)))
			}
		}
	}
	return newColor
}
func imfilterMid(filter [][]float32, rect image.Rectangle, rgbFunc func(xpos, ypos int) (r0, g0, b0, a0 uint8)) *image.RGBA {
	size := rect.Size()
	imgGrey := image.NewRGBA(rect)
	fX := len(filter)
	fY := len(filter[0])
	for x := fX / 2; x < size.X-fX/2; x++ {
		for y := fY / 2; y < size.Y-fY/2; y++ {
			_, _, _, a := rgbFunc(x, y)
			newColor := color.RGBA{R: 0, G: 0, B: 0, A: a}
			imgGrey.SetRGBA(x, y, newColor)
		}
	}
	return imgGrey
}
