package easygif

import (
	"image"
	"image/png"
	"os"
)

func SaveImageToPNG(img image.Image, outputPNGPath string) error {
	out, err := os.Create(outputPNGPath)
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, img)
	if err != nil {
		return err
	}
	return nil
}
