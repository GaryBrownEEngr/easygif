package main

import "easygif"

func main() {
	img, err := easygif.Screenshot()
	if err != nil {
		panic(err)
	}

	err = easygif.SaveImageToPNG(img, "./examples/screenshot/screenshot.png")
	if err != nil {
		panic(err)
	}

	// trimmed
	img, err = easygif.ScreenshotTrimmed(200, 200, 200, 600)
	if err != nil {
		panic(err)
	}

	err = easygif.SaveImageToPNG(img, "./examples/screenshot/screenshotTrimmed.png")
	if err != nil {
		panic(err)
	}
}
