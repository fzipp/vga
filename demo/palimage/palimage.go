// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"image"
	"image/png"
	"log"
	"os"

	"github.com/fzipp/vga"
)

const imageFileName = "vga256_palette.png"

func main() {
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
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()
	err = png.Encode(outFile, img)
	if err != nil {
		log.Fatal(err)
	}
}

func drawRectangle(img *image.Paletted, colorIndex uint8, x, y, w, h int) {
	for ix := x; ix < x+w; ix++ {
		for iy := y; iy < y+h; iy++ {
			img.SetColorIndex(ix, iy, colorIndex)
		}
	}
}
