package main

import (
	"time"

	"github.com/GaryBrownEEngr/easygif"
)

func main() {
	frames, _ := easygif.ScreenshotVideo(50, time.Millisecond*100)
	_ = easygif.EasyGifWrite(frames, time.Millisecond*100, "./examples/gif/screenshot.gif")

	frames, _ = easygif.ScreenshotVideoTrimmed(50, time.Millisecond*100, 200, 10, 50, 400)
	_ = easygif.EasyGifWrite(frames, time.Millisecond*100, "./examples/gif/screenshotTrimmed.gif")
}
