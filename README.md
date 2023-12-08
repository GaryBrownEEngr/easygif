# EasyGif

Easy creation of GIF images with Golang

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

```go
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

```