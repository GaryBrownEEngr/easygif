package main

import (
	"easygif"
	"time"
)

func main() {
	frames, err := easygif.ScreenshotVideoTrimmed(50, time.Millisecond*100, 200, 10, 50, 400)
	if err != nil {
		panic(err)
	}

	err = easygif.EasyGifWrite(frames, time.Millisecond*100, "./examples/gif/screenshotTrimmed.gif")
	if err != nil {
		panic(err)
	}
}
