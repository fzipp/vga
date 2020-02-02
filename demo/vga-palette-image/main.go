// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A little demo program that creates a PNG image of the VGA 256 color
// default palette.
//
// Usage:
//     vga-palette-image vga256_palette.png
package main

import (
	"fmt"
	"image"
	"image/png"
	"os"

	"github.com/fzipp/vga"
)

func usage() {
	fail(`A little demo program that creates a PNG image of the VGA 256 color
default palette.

Usage: vga-palette-image vga256_palette.png`)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	imageFileName := os.Args[1]

	colorCount := 256
	imageWidth := 256
	imageHeight := imageWidth
	colorsPerRow := 16
	colorBoxWidth := imageWidth / colorsPerRow
	rows := colorCount / colorsPerRow
	colorBoxHeight := imageHeight / rows
	margin := 1

	img := image.NewPaletted(image.Rect(0, 0, imageWidth, imageHeight),
		vga.DefaultPalette)

	for colorIndex := 0; colorIndex < colorCount; colorIndex++ {
		row := colorIndex / colorsPerRow
		column := colorIndex % colorsPerRow
		x := column*colorBoxWidth + margin
		y := row*colorBoxHeight + margin
		width := colorBoxWidth - 2*margin
		height := colorBoxHeight - 2*margin

		drawRectangle(img, uint8(colorIndex), x, y, width, height)
	}

	outFile, err := os.Create(imageFileName)
	check(err)
	defer outFile.Close()
	err = png.Encode(outFile, img)
	check(err)
}

func drawRectangle(img *image.Paletted, colorIndex uint8, x, y, w, h int) {
	for ix := x; ix < x+w; ix++ {
		for iy := y; iy < y+h; iy++ {
			img.SetColorIndex(ix, iy, colorIndex)
		}
	}
}

func check(err error) {
	if err != nil {
		fail(err)
	}
}

func fail(message interface{}) {
	_, _ = fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
