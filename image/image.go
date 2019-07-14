package image


import (
	"image"
	"image/color"
	U "github.com/neurons-platform/gotools/utils"
	"os"
)


// Create a struct to deal with pixel
type Pixel struct {
	Point image.Point
	Color color.Color
}

// Keep it DRY so don't have to repeat opening file and decode
func OpenAndDecode(filepath string) (image.Image, string, error) {
	imgFile, err := os.Open(filepath)
	if err != nil {
		// panic(err)
		U.Throw(err)
	}

	defer imgFile.Close()
	img, format, err := image.Decode(imgFile)
	if err != nil {
		// panic(err)
		U.Throw(err)
	}

	return img, format, nil
}

// Decode image.Image's pixel data into []*Pixel
func DecodePixelsFromImage(img image.Image, offsetX, offsetY int) []*Pixel {
	pixels := []*Pixel{}
	for y := 0; y <= img.Bounds().Max.Y; y++ {
		for x := 0; x <= img.Bounds().Max.X; x++ {
			p := &Pixel{
				Point: image.Point{x + offsetX, y + offsetY},
				Color: img.At(x, y),
			}
			pixels = append(pixels, p)
		}
	}
	return pixels
}
