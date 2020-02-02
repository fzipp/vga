# vga

The VGA 256-color [default palette](https://en.wikipedia.org/wiki/Video_Graphics_Array#Color_palette)
for Go. It can be used as a [color.Palette](https://golang.org/pkg/image/color/#Palette)
from the standard library, e.g to create an [image.Paletted](https://golang.org/pkg/image/#Paletted).

    go get github.com/fzipp/vga

## Example usage

    package main

    import (
    	"image"

    	"github.com/fzipp/vga"
    )

    func main() {
    	img := image.NewPaletted(image.Rect(0, 0, 320, 200), vga.DefaultPalette)
    	// ...
    }

