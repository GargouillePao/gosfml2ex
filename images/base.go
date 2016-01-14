package images

import (
	"errors"
	"image"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"strings"
)

func DecodeImage(fileName string) (img image.Image, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	suffixName := strings.ToLower(fileName[strings.Index(fileName, ".")+1:])
	switch suffixName {
	case "jpg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	default:
		err = errors.New("Not correct image format :" + suffixName)
	}
	return
}
func ReadImage(fileName string, x, y, width, height int) (imageGot *image.RGBA, point image.Point, err error) {

	imgSrc, err := DecodeImage(fileName)

	rect := image.Rect(x, y, x+width, y+height)
	imageGot = image.NewRGBA(rect)
	point = image.Point{x, y}
	draw.Draw(imageGot, imageGot.Bounds(), imgSrc, point, draw.Src)
	return
}
func WriteImage(fileName string, img image.RGBA) error {
	suffixName := strings.ToLower(fileName[strings.Index(fileName, ".")+1:])
	file, err := os.OpenFile(fileName, os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	switch suffixName {
	case "jpg":
		jpeg.Encode(file, &img, &jpeg.Options{90})
	case "png":
		png.Encode(file, &img)
	case "gif":
		gif.Encode(file, &img, nil)
	default:
		err = errors.New("Not correct image format :" + suffixName)
	}
	if err != nil {
		return err
	}
	return nil
}
