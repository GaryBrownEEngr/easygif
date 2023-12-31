# EasyGif

Easy creation of screenshots and GIFs with Golang.

[![Go Reference](https://pkg.go.dev/badge/github.com/GaryBrownEEngr/easygif.svg)](https://pkg.go.dev/github.com/GaryBrownEEngr/easygif)
[![Go CI](https://github.com/GaryBrownEEngr/easygif/actions/workflows/go.yml/badge.svg)](https://github.com/GaryBrownEEngr/easygif/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/GaryBrownEEngr/easygif)](https://goreportcard.com/report/github.com/GaryBrownEEngr/easygif)
[![Coverage Badge](https://img.shields.io/endpoint?url=https://gist.githubusercontent.com/GaryBrownEEngr/0a036dc69ea9afb4202e2d262fec1e1d/raw/GaryBrownEEngr_easygif_main.json)](https://github.com/GaryBrownEEngr/easygif/actions)

## Install

```bash
go get github.com/GaryBrownEEngr/easygif
```

## Easy Screenshots

```go
package main

import "github.com/GaryBrownEEngr/easygif"

func main() {
	img, _ := easygif.Screenshot()
	_ = easygif.SaveImageToPNG(img, "./examples/screenshot/screenshot.png")

	// trimmed
	img, _ = easygif.ScreenshotTrimmed(200, 200, 200, 600)
	_ = easygif.SaveImageToPNG(img, "./examples/screenshot/screenshotTrimmed.png")
}

```

## Easy GIF Creation

This package layers on top of the built-in golang [image/gif](https://pkg.go.dev/image/gif) package. The goal is to remove all the complexity of making a GIF with go. Once you have a slice of images then only a single function call is required to generate and write your GIF.

```go
package main

import (
	"time"

	"github.com/GaryBrownEEngr/easygif"
)

func main() {
	// Collect screenshots of either the entire screen or a trimmed section of it.
	//frames, _ := easygif.ScreenshotVideo(50, time.Millisecond*100)
	frames, _ := easygif.ScreenshotVideoTrimmed(50, time.Millisecond*100, 200, 10, 50, 400)

	// Create a GIF using the Plan9 color palette and nearest color approximation.
	easygif.NearestWrite(frames, time.Millisecond*100, "easy.gif")

	// Use dithering for better colors, but adds noise
	easygif.DitheredWrite(frames, time.Millisecond*100, "easyDithered.gif")

	// Use the 256 most common colors found in the frames
	easygif.MostCommonColorsWrite(frames, time.Millisecond*100, "mostCommonColors.gif")
}

```

### easygif.Nearest

![GIF made by golang easygif - One does not simply make a gif - using nearest Plan9 color](https://github.com/GaryBrownEEngr/easygif/blob/main/examples/gif/OneDoesNotSimplyMakeAGIF_Nearest.gif)

### easygif.Dithered

![GIF made by golang easygif - One does not simply make a gif - using dithering](https://github.com/GaryBrownEEngr/easygif/blob/main/examples/gif/OneDoesNotSimplyMakeAGIF_Dithered.gif)

### easygif.MostCommonColors

![GIF made by golang easygif - One does not simply make a gif - using the 256 most common colors](https://github.com/GaryBrownEEngr/easygif/blob/main/examples/gif/OneDoesNotSimplyMakeAGIF_MostCommon.gif)

Screenshots taken from [here](https://www.youtube.com/watch?v=klidgum0_v8).

I would recommend using `easygif.MostCommonColors`. Happy GIFing.
