// Copyright 2020 Frederik Zipp. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vga provides the VGA 256-color default palette, famously used in
// video mode 13h.
// See also https://en.wikipedia.org/wiki/Video_Graphics_Array#Color_palette
package vga

import "image/color"

// DefaultPalette is the VGA 256-color default palette. It can be used as a
// color.Palette from the standard library, e.g to create an image.Paletted.
var DefaultPalette = []color.Color{
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0xaa, A: 0xff},
	color.RGBA{R: 0x00, G: 0xaa, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0xaa, B: 0xaa, A: 0xff},
	color.RGBA{R: 0xaa, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0xaa, G: 0x00, B: 0xaa, A: 0xff},
	color.RGBA{R: 0xaa, G: 0x55, B: 0x00, A: 0xff},
	color.RGBA{R: 0xaa, G: 0xaa, B: 0xaa, A: 0xff},
	color.RGBA{R: 0x55, G: 0x55, B: 0x55, A: 0xff},
	color.RGBA{R: 0x55, G: 0x55, B: 0xff, A: 0xff},
	color.RGBA{R: 0x55, G: 0xff, B: 0x55, A: 0xff},
	color.RGBA{R: 0x55, G: 0xff, B: 0xff, A: 0xff},
	color.RGBA{R: 0xff, G: 0x55, B: 0x55, A: 0xff},
	color.RGBA{R: 0xff, G: 0x55, B: 0xff, A: 0xff},
	color.RGBA{R: 0xff, G: 0xff, B: 0x55, A: 0xff},
	color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x14, G: 0x14, B: 0x14, A: 0xff},
	color.RGBA{R: 0x20, G: 0x20, B: 0x20, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x2c, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x38, G: 0x38, B: 0x38, A: 0xff},
	color.RGBA{R: 0x45, G: 0x45, B: 0x45, A: 0xff},
	color.RGBA{R: 0x51, G: 0x51, B: 0x51, A: 0xff},
	color.RGBA{R: 0x61, G: 0x61, B: 0x61, A: 0xff},
	color.RGBA{R: 0x71, G: 0x71, B: 0x71, A: 0xff},
	color.RGBA{R: 0x82, G: 0x82, B: 0x82, A: 0xff},
	color.RGBA{R: 0x92, G: 0x92, B: 0x92, A: 0xff},
	color.RGBA{R: 0xa2, G: 0xa2, B: 0xa2, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xb6, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xcb, G: 0xcb, B: 0xcb, A: 0xff},
	color.RGBA{R: 0xe3, G: 0xe3, B: 0xe3, A: 0xff},
	color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff},
	color.RGBA{R: 0x41, G: 0x00, B: 0xff, A: 0xff},
	color.RGBA{R: 0x7d, G: 0x00, B: 0xff, A: 0xff},
	color.RGBA{R: 0xbe, G: 0x00, B: 0xff, A: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0xbe, A: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0x7d, A: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0x41, A: 0xff},
	color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0xff, G: 0x41, B: 0x00, A: 0xff},
	color.RGBA{R: 0xff, G: 0x7d, B: 0x00, A: 0xff},
	color.RGBA{R: 0xff, G: 0xbe, B: 0x00, A: 0xff},
	color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xff},
	color.RGBA{R: 0xbe, G: 0xff, B: 0x00, A: 0xff},
	color.RGBA{R: 0x7d, G: 0xff, B: 0x00, A: 0xff},
	color.RGBA{R: 0x41, G: 0xff, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0xff, B: 0x41, A: 0xff},
	color.RGBA{R: 0x00, G: 0xff, B: 0x7d, A: 0xff},
	color.RGBA{R: 0x00, G: 0xff, B: 0xbe, A: 0xff},
	color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff},
	color.RGBA{R: 0x00, G: 0xbe, B: 0xff, A: 0xff},
	color.RGBA{R: 0x00, G: 0x7d, B: 0xff, A: 0xff},
	color.RGBA{R: 0x00, G: 0x41, B: 0xff, A: 0xff},
	color.RGBA{R: 0x7d, G: 0x7d, B: 0xff, A: 0xff},
	color.RGBA{R: 0x9e, G: 0x7d, B: 0xff, A: 0xff},
	color.RGBA{R: 0xbe, G: 0x7d, B: 0xff, A: 0xff},
	color.RGBA{R: 0xdf, G: 0x7d, B: 0xff, A: 0xff},
	color.RGBA{R: 0xff, G: 0x7d, B: 0xff, A: 0xff},
	color.RGBA{R: 0xff, G: 0x7d, B: 0xdf, A: 0xff},
	color.RGBA{R: 0xff, G: 0x7d, B: 0xbe, A: 0xff},
	color.RGBA{R: 0xff, G: 0x7d, B: 0x9e, A: 0xff},
	color.RGBA{R: 0xff, G: 0x7d, B: 0x7d, A: 0xff},
	color.RGBA{R: 0xff, G: 0x9e, B: 0x7d, A: 0xff},
	color.RGBA{R: 0xff, G: 0xbe, B: 0x7d, A: 0xff},
	color.RGBA{R: 0xff, G: 0xdf, B: 0x7d, A: 0xff},
	color.RGBA{R: 0xff, G: 0xff, B: 0x7d, A: 0xff},
	color.RGBA{R: 0xdf, G: 0xff, B: 0x7d, A: 0xff},
	color.RGBA{R: 0xbe, G: 0xff, B: 0x7d, A: 0xff},
	color.RGBA{R: 0x9e, G: 0xff, B: 0x7d, A: 0xff},
	color.RGBA{R: 0x7d, G: 0xff, B: 0x7d, A: 0xff},
	color.RGBA{R: 0x7d, G: 0xff, B: 0x9e, A: 0xff},
	color.RGBA{R: 0x7d, G: 0xff, B: 0xbe, A: 0xff},
	color.RGBA{R: 0x7d, G: 0xff, B: 0xdf, A: 0xff},
	color.RGBA{R: 0x7d, G: 0xff, B: 0xff, A: 0xff},
	color.RGBA{R: 0x7d, G: 0xdf, B: 0xff, A: 0xff},
	color.RGBA{R: 0x7d, G: 0xbe, B: 0xff, A: 0xff},
	color.RGBA{R: 0x7d, G: 0x9e, B: 0xff, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xb6, B: 0xff, A: 0xff},
	color.RGBA{R: 0xc7, G: 0xb6, B: 0xff, A: 0xff},
	color.RGBA{R: 0xdb, G: 0xb6, B: 0xff, A: 0xff},
	color.RGBA{R: 0xeb, G: 0xb6, B: 0xff, A: 0xff},
	color.RGBA{R: 0xff, G: 0xb6, B: 0xff, A: 0xff},
	color.RGBA{R: 0xff, G: 0xb6, B: 0xeb, A: 0xff},
	color.RGBA{R: 0xff, G: 0xb6, B: 0xdb, A: 0xff},
	color.RGBA{R: 0xff, G: 0xb6, B: 0xc7, A: 0xff},
	color.RGBA{R: 0xff, G: 0xb6, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xff, G: 0xc7, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xff, G: 0xdb, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xff, G: 0xeb, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xff, G: 0xff, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xeb, G: 0xff, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xdb, G: 0xff, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xc7, G: 0xff, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xff, B: 0xb6, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xff, B: 0xc7, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xff, B: 0xdb, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xff, B: 0xeb, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xff, B: 0xff, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xeb, B: 0xff, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xdb, B: 0xff, A: 0xff},
	color.RGBA{R: 0xb6, G: 0xc7, B: 0xff, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x71, A: 0xff},
	color.RGBA{R: 0x1c, G: 0x00, B: 0x71, A: 0xff},
	color.RGBA{R: 0x38, G: 0x00, B: 0x71, A: 0xff},
	color.RGBA{R: 0x55, G: 0x00, B: 0x71, A: 0xff},
	color.RGBA{R: 0x71, G: 0x00, B: 0x71, A: 0xff},
	color.RGBA{R: 0x71, G: 0x00, B: 0x55, A: 0xff},
	color.RGBA{R: 0x71, G: 0x00, B: 0x38, A: 0xff},
	color.RGBA{R: 0x71, G: 0x00, B: 0x1c, A: 0xff},
	color.RGBA{R: 0x71, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x71, G: 0x1c, B: 0x00, A: 0xff},
	color.RGBA{R: 0x71, G: 0x38, B: 0x00, A: 0xff},
	color.RGBA{R: 0x71, G: 0x55, B: 0x00, A: 0xff},
	color.RGBA{R: 0x71, G: 0x71, B: 0x00, A: 0xff},
	color.RGBA{R: 0x55, G: 0x71, B: 0x00, A: 0xff},
	color.RGBA{R: 0x38, G: 0x71, B: 0x00, A: 0xff},
	color.RGBA{R: 0x1c, G: 0x71, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x71, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x71, B: 0x1c, A: 0xff},
	color.RGBA{R: 0x00, G: 0x71, B: 0x38, A: 0xff},
	color.RGBA{R: 0x00, G: 0x71, B: 0x55, A: 0xff},
	color.RGBA{R: 0x00, G: 0x71, B: 0x71, A: 0xff},
	color.RGBA{R: 0x00, G: 0x55, B: 0x71, A: 0xff},
	color.RGBA{R: 0x00, G: 0x38, B: 0x71, A: 0xff},
	color.RGBA{R: 0x00, G: 0x1c, B: 0x71, A: 0xff},
	color.RGBA{R: 0x38, G: 0x38, B: 0x71, A: 0xff},
	color.RGBA{R: 0x45, G: 0x38, B: 0x71, A: 0xff},
	color.RGBA{R: 0x55, G: 0x38, B: 0x71, A: 0xff},
	color.RGBA{R: 0x61, G: 0x38, B: 0x71, A: 0xff},
	color.RGBA{R: 0x71, G: 0x38, B: 0x71, A: 0xff},
	color.RGBA{R: 0x71, G: 0x38, B: 0x61, A: 0xff},
	color.RGBA{R: 0x71, G: 0x38, B: 0x55, A: 0xff},
	color.RGBA{R: 0x71, G: 0x38, B: 0x45, A: 0xff},
	color.RGBA{R: 0x71, G: 0x38, B: 0x38, A: 0xff},
	color.RGBA{R: 0x71, G: 0x45, B: 0x38, A: 0xff},
	color.RGBA{R: 0x71, G: 0x55, B: 0x38, A: 0xff},
	color.RGBA{R: 0x71, G: 0x61, B: 0x38, A: 0xff},
	color.RGBA{R: 0x71, G: 0x71, B: 0x38, A: 0xff},
	color.RGBA{R: 0x61, G: 0x71, B: 0x38, A: 0xff},
	color.RGBA{R: 0x55, G: 0x71, B: 0x38, A: 0xff},
	color.RGBA{R: 0x45, G: 0x71, B: 0x38, A: 0xff},
	color.RGBA{R: 0x38, G: 0x71, B: 0x38, A: 0xff},
	color.RGBA{R: 0x38, G: 0x71, B: 0x45, A: 0xff},
	color.RGBA{R: 0x38, G: 0x71, B: 0x55, A: 0xff},
	color.RGBA{R: 0x38, G: 0x71, B: 0x61, A: 0xff},
	color.RGBA{R: 0x38, G: 0x71, B: 0x71, A: 0xff},
	color.RGBA{R: 0x38, G: 0x61, B: 0x71, A: 0xff},
	color.RGBA{R: 0x38, G: 0x55, B: 0x71, A: 0xff},
	color.RGBA{R: 0x38, G: 0x45, B: 0x71, A: 0xff},
	color.RGBA{R: 0x51, G: 0x51, B: 0x71, A: 0xff},
	color.RGBA{R: 0x59, G: 0x51, B: 0x71, A: 0xff},
	color.RGBA{R: 0x61, G: 0x51, B: 0x71, A: 0xff},
	color.RGBA{R: 0x69, G: 0x51, B: 0x71, A: 0xff},
	color.RGBA{R: 0x71, G: 0x51, B: 0x71, A: 0xff},
	color.RGBA{R: 0x71, G: 0x51, B: 0x69, A: 0xff},
	color.RGBA{R: 0x71, G: 0x51, B: 0x61, A: 0xff},
	color.RGBA{R: 0x71, G: 0x51, B: 0x59, A: 0xff},
	color.RGBA{R: 0x71, G: 0x51, B: 0x51, A: 0xff},
	color.RGBA{R: 0x71, G: 0x59, B: 0x51, A: 0xff},
	color.RGBA{R: 0x71, G: 0x61, B: 0x51, A: 0xff},
	color.RGBA{R: 0x71, G: 0x69, B: 0x51, A: 0xff},
	color.RGBA{R: 0x71, G: 0x71, B: 0x51, A: 0xff},
	color.RGBA{R: 0x69, G: 0x71, B: 0x51, A: 0xff},
	color.RGBA{R: 0x61, G: 0x71, B: 0x51, A: 0xff},
	color.RGBA{R: 0x59, G: 0x71, B: 0x51, A: 0xff},
	color.RGBA{R: 0x51, G: 0x71, B: 0x51, A: 0xff},
	color.RGBA{R: 0x51, G: 0x71, B: 0x59, A: 0xff},
	color.RGBA{R: 0x51, G: 0x71, B: 0x61, A: 0xff},
	color.RGBA{R: 0x51, G: 0x71, B: 0x69, A: 0xff},
	color.RGBA{R: 0x51, G: 0x71, B: 0x71, A: 0xff},
	color.RGBA{R: 0x51, G: 0x69, B: 0x71, A: 0xff},
	color.RGBA{R: 0x51, G: 0x61, B: 0x71, A: 0xff},
	color.RGBA{R: 0x51, G: 0x59, B: 0x71, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x41, A: 0xff},
	color.RGBA{R: 0x10, G: 0x00, B: 0x41, A: 0xff},
	color.RGBA{R: 0x20, G: 0x00, B: 0x41, A: 0xff},
	color.RGBA{R: 0x30, G: 0x00, B: 0x41, A: 0xff},
	color.RGBA{R: 0x41, G: 0x00, B: 0x41, A: 0xff},
	color.RGBA{R: 0x41, G: 0x00, B: 0x30, A: 0xff},
	color.RGBA{R: 0x41, G: 0x00, B: 0x20, A: 0xff},
	color.RGBA{R: 0x41, G: 0x00, B: 0x10, A: 0xff},
	color.RGBA{R: 0x41, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x41, G: 0x10, B: 0x00, A: 0xff},
	color.RGBA{R: 0x41, G: 0x20, B: 0x00, A: 0xff},
	color.RGBA{R: 0x41, G: 0x30, B: 0x00, A: 0xff},
	color.RGBA{R: 0x41, G: 0x41, B: 0x00, A: 0xff},
	color.RGBA{R: 0x30, G: 0x41, B: 0x00, A: 0xff},
	color.RGBA{R: 0x20, G: 0x41, B: 0x00, A: 0xff},
	color.RGBA{R: 0x10, G: 0x41, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x41, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x41, B: 0x10, A: 0xff},
	color.RGBA{R: 0x00, G: 0x41, B: 0x20, A: 0xff},
	color.RGBA{R: 0x00, G: 0x41, B: 0x30, A: 0xff},
	color.RGBA{R: 0x00, G: 0x41, B: 0x41, A: 0xff},
	color.RGBA{R: 0x00, G: 0x30, B: 0x41, A: 0xff},
	color.RGBA{R: 0x00, G: 0x20, B: 0x41, A: 0xff},
	color.RGBA{R: 0x00, G: 0x10, B: 0x41, A: 0xff},
	color.RGBA{R: 0x20, G: 0x20, B: 0x41, A: 0xff},
	color.RGBA{R: 0x28, G: 0x20, B: 0x41, A: 0xff},
	color.RGBA{R: 0x30, G: 0x20, B: 0x41, A: 0xff},
	color.RGBA{R: 0x38, G: 0x20, B: 0x41, A: 0xff},
	color.RGBA{R: 0x41, G: 0x20, B: 0x41, A: 0xff},
	color.RGBA{R: 0x41, G: 0x20, B: 0x38, A: 0xff},
	color.RGBA{R: 0x41, G: 0x20, B: 0x30, A: 0xff},
	color.RGBA{R: 0x41, G: 0x20, B: 0x28, A: 0xff},
	color.RGBA{R: 0x41, G: 0x20, B: 0x20, A: 0xff},
	color.RGBA{R: 0x41, G: 0x28, B: 0x20, A: 0xff},
	color.RGBA{R: 0x41, G: 0x30, B: 0x20, A: 0xff},
	color.RGBA{R: 0x41, G: 0x38, B: 0x20, A: 0xff},
	color.RGBA{R: 0x41, G: 0x41, B: 0x20, A: 0xff},
	color.RGBA{R: 0x38, G: 0x41, B: 0x20, A: 0xff},
	color.RGBA{R: 0x30, G: 0x41, B: 0x20, A: 0xff},
	color.RGBA{R: 0x28, G: 0x41, B: 0x20, A: 0xff},
	color.RGBA{R: 0x20, G: 0x41, B: 0x20, A: 0xff},
	color.RGBA{R: 0x20, G: 0x41, B: 0x28, A: 0xff},
	color.RGBA{R: 0x20, G: 0x41, B: 0x30, A: 0xff},
	color.RGBA{R: 0x20, G: 0x41, B: 0x38, A: 0xff},
	color.RGBA{R: 0x20, G: 0x41, B: 0x41, A: 0xff},
	color.RGBA{R: 0x20, G: 0x38, B: 0x41, A: 0xff},
	color.RGBA{R: 0x20, G: 0x30, B: 0x41, A: 0xff},
	color.RGBA{R: 0x20, G: 0x28, B: 0x41, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x2c, B: 0x41, A: 0xff},
	color.RGBA{R: 0x30, G: 0x2c, B: 0x41, A: 0xff},
	color.RGBA{R: 0x34, G: 0x2c, B: 0x41, A: 0xff},
	color.RGBA{R: 0x3c, G: 0x2c, B: 0x41, A: 0xff},
	color.RGBA{R: 0x41, G: 0x2c, B: 0x41, A: 0xff},
	color.RGBA{R: 0x41, G: 0x2c, B: 0x3c, A: 0xff},
	color.RGBA{R: 0x41, G: 0x2c, B: 0x34, A: 0xff},
	color.RGBA{R: 0x41, G: 0x2c, B: 0x30, A: 0xff},
	color.RGBA{R: 0x41, G: 0x2c, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x41, G: 0x30, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x41, G: 0x34, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x41, G: 0x3c, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x41, G: 0x41, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x3c, G: 0x41, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x34, G: 0x41, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x30, G: 0x41, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x41, B: 0x2c, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x41, B: 0x30, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x41, B: 0x34, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x41, B: 0x3c, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x41, B: 0x41, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x3c, B: 0x41, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x34, B: 0x41, A: 0xff},
	color.RGBA{R: 0x2c, G: 0x30, B: 0x41, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
	color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xff},
}
