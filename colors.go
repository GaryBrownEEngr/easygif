package easygif

import "image/color"

// https://www.w3schools.com/colors/colors_names.asp
var (
	// Black to white
	Black     color.RGBA = color.RGBA{0x00, 0x00, 0x00, 0xFF} // #000000
	DarkGray  color.RGBA = color.RGBA{0x26, 0x26, 0x26, 0xFF} // #262626
	Gray      color.RGBA = color.RGBA{0x80, 0x80, 0x80, 0xFF} // #808080
	LightGray color.RGBA = color.RGBA{0xD3, 0xD3, 0xD3, 0xFF} // #D3D3D3
	White     color.RGBA = color.RGBA{0xFF, 0xFF, 0xFF, 0xFF} // #FFFFFF

	// Primary Colors
	Red  color.RGBA = color.RGBA{0xFF, 0x00, 0x00, 0xFF} // #FF0000
	Lime color.RGBA = color.RGBA{0x00, 0xFF, 0x00, 0xFF} // #00FF00
	Blue color.RGBA = color.RGBA{0x00, 0x00, 0xFF, 0xFF} // #0000FF

	// half strength primary colors
	Maroon   color.RGBA = color.RGBA{0x80, 0x00, 0x00, 0xFF} // #800000
	Green    color.RGBA = color.RGBA{0x00, 0x80, 0x00, 0xFF} // #008000
	NavyBlue color.RGBA = color.RGBA{0x00, 0x00, 0x80, 0xFF} // #000080

	// full strength primary mixes
	Yellow  color.RGBA = color.RGBA{0xFF, 0xFF, 0x00, 0xFF} // #FFFF00
	Aqua    color.RGBA = color.RGBA{0x00, 0xFF, 0xFF, 0xFF} // #00FFFF
	Cyan               = Aqua                               // #00FFFF
	Magenta color.RGBA = color.RGBA{0xFF, 0x00, 0xFF, 0xFF} // #FF00FF
	Fuchsia            = Magenta                            // #FF00FF

	// half strength primary mixes
	Olive  color.RGBA = color.RGBA{0x80, 0x80, 0x00, 0xFF} // #808000
	Purple color.RGBA = color.RGBA{0x80, 0x00, 0x80, 0xFF} // #800080
	Teal   color.RGBA = color.RGBA{0x00, 0x80, 0x80, 0xFF} // #008080

	// Other interesting colors
	Orange      color.RGBA = color.RGBA{0xFF, 0xA5, 0x00, 0xFF} // #FFA500
	Indigo      color.RGBA = color.RGBA{0x4B, 0x00, 0x82, 0xFF} // #4B0082
	Violet      color.RGBA = color.RGBA{0xEE, 0x82, 0xEE, 0xFF} // #EE82EE
	Gold        color.RGBA = color.RGBA{0xFF, 0xD7, 0x00, 0xFF} // #FFD700
	SkyBlue     color.RGBA = color.RGBA{0x87, 0xCE, 0xEB, 0xFF} // #87CEEB
	SaddleBrown color.RGBA = color.RGBA{0x8B, 0x45, 0x13, 0xFF} // #8B4513
	Tan         color.RGBA = color.RGBA{0xD2, 0xB4, 0x8C, 0xFF} // #D2B48C
	Crimson     color.RGBA = color.RGBA{0xDC, 0x14, 0x3C, 0xFF} // #DC143C
	Pink        color.RGBA = color.RGBA{0xFF, 0xC0, 0xCB, 0xFF} // #FFC0CB
	Salmon      color.RGBA = color.RGBA{0xFA, 0x80, 0x72, 0xFF} // #FA8072
	Turquoise   color.RGBA = color.RGBA{0x40, 0xE0, 0xD0, 0xFF} // #40E0D0
)

// Creates a color between the given a and b. 0 means a is given, 1 means b is given, .5 is a color half way between.
// The ratio is capped between 0 and 1
// Currently the function floors the number instead of rounding to nearest.
func LerpColor(a, b color.RGBA, ratio float64) color.RGBA {
	ret := color.RGBA{
		R: Lerp(a.R, b.R, ratio),
		G: Lerp(a.G, b.G, ratio),
		B: Lerp(a.B, b.B, ratio),
		A: Lerp(a.A, b.A, ratio),
	}

	return ret
}
