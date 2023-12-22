package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"time"

	"github.com/GaryBrownEEngr/easygif"
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	chooseNearest()
	useDithering()
	useMostCommonColors()
	easyGifExample()
	// EasyGifExample2()
}

func chooseNearest() {
	fileData, _ := os.ReadFile("OneDoesNotSimply_Template.jpg")
	img, _ := jpeg.Decode(bytes.NewReader(fileData))

	bound := img.Bounds()

	palettedImg := image.NewPaletted(bound, palette.Plan9)

	startTime := time.Now()
	draw.Draw(palettedImg, bound, img, image.Point{}, draw.Src)
	fmt.Println(time.Since(startTime))

	anim := gif.GIF{}
	anim.Image = append(anim.Image, palettedImg)
	anim.Delay = append(anim.Delay, 100)

	file, _ := os.Create("OneDoesNotSimply_ChooseNearestColor.gif")
	defer file.Close()

	_ = gif.EncodeAll(file, &anim)
}

func useDithering() {
	fileData, _ := os.ReadFile("OneDoesNotSimply_Template.jpg")
	img, _ := jpeg.Decode(bytes.NewReader(fileData))

	bound := img.Bounds()
	palettedImg := image.NewPaletted(bound, palette.Plan9)
	drawer := draw.FloydSteinberg

	startTime := time.Now()
	drawer.Draw(palettedImg, bound, img, image.Point{})
	fmt.Println(time.Since(startTime))

	anim := gif.GIF{}
	anim.Image = append(anim.Image, palettedImg)
	anim.Delay = append(anim.Delay, 100)

	file, _ := os.Create("OneDoesNotSimply_UseDithering.gif")
	defer file.Close()

	_ = gif.EncodeAll(file, &anim)
}

func useMostCommonColors() {
	fileData, _ := os.ReadFile("OneDoesNotSimply_Template.jpg")
	img, _ := jpeg.Decode(bytes.NewReader(fileData))
	startTime := time.Now()
	_ = easygif.MostCommonColorsGifWrite([]image.Image{img}, time.Second, "OneDoesNotSimply_UseMostCommonColors.gif")
	fmt.Println(time.Since(startTime))
}

func addTextToCenterOfImage(img image.Image, text string, c color.Color, fontSize float64) image.Image {
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: fontSize,
	})

	dc := gg.NewContextForImage(img)
	bound := img.Bounds()
	dc.SetFontFace(face)
	dc.SetColor(c)
	dc.DrawStringAnchored(text, float64(bound.Dx())/2, float64(bound.Dy())*.45, 0.5, 0.5)
	// dc.DrawStringAnchored(text, 0, 0, 0.5, 0.5)
	return dc.Image()
}

func easyGifExample() {
	var img image.Image = image.NewRGBA(image.Rect(0, 0, 100, 100))
	imgA := addTextToCenterOfImage(img, "A", easygif.Red, 60)
	imgB := addTextToCenterOfImage(img, "B", easygif.Green, 60)
	imgC := addTextToCenterOfImage(img, "C", easygif.Blue, 60)
	_ = easygif.SaveImageToPNG(imgA, "./imageDirectory/A.png")
	_ = easygif.SaveImageToPNG(imgB, "./imageDirectory/B.png")
	_ = easygif.SaveImageToPNG(imgC, "./imageDirectory/C.png")

	imageDirectory := "./imageDirectory"
	files, _ := os.ReadDir(imageDirectory)
	frames := []image.Image{}
	for _, file := range files {
		fileData, _ := os.ReadFile(path.Join(imageDirectory, file.Name()))
		img, _ := png.Decode(bytes.NewReader(fileData))
		frames = append(frames, img)
	}
	_ = easygif.EasyGifWrite(frames, time.Millisecond*1000, "easyGif.gif")
}

func EasyGifExample2() {
	frames, _ := easygif.ScreenshotVideo(30, time.Millisecond*100)
	_ = easygif.EasyDitheredGifWrite(frames, time.Millisecond*100, "easyGif.gif")
}
