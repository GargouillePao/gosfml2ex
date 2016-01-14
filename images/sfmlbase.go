package images

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"image"
)

func ReadTexture(fileName string, x, y, width, height int) (texture *sf.Texture, point image.Point, err error) {
	imageGot, point, err := ReadImage(fileName, x, y, width, height)
	if err != nil {
		return
	}
	imagesf, err := sf.NewImageFromPixels(uint(width), uint(height), imageGot.Pix)
	if err != nil {
		return
	}
	texture, err = sf.NewTextureFromImage(imagesf, &sf.IntRect{0, 0, width, height})
	if err != nil {
		return
	}
	return
}
func ReadTextureFromImage(imageGot image.RGBA) (texture *sf.Texture, err error) {
	rect := imageGot.Rect

	imagesf, err := sf.NewImageFromPixels(uint(rect.Size().X), uint(rect.Size().Y), imageGot.Pix)
	if err != nil {
		return
	}
	texture, err = sf.NewTextureFromImage(imagesf, &sf.IntRect{0, 0, rect.Size().X, rect.Size().Y})
	if err != nil {
		return
	}
	return
}
