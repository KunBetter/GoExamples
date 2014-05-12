// pic
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

type img struct {
	h, w int
	m    [][]color.RGBA
}

func (m *img) At(x, y int) color.Color { return m.m[x][y] }
func (m *img) ColorModel() color.Model { return color.RGBAModel }
func (m *img) Bounds() image.Rectangle { return image.Rect(0, 0, m.h, m.w) }

func drawPixels(h, w int) image.Image {
	c := make([][]color.RGBA, h)
	for i := range c {
		c[i] = make([]color.RGBA, w)
	}

	m := &img{h, w, c}

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c[y][x].R = byte(x * y)
			c[y][x].G = byte(x * y * 2)
			c[y][x].B = byte(x * y * 3)
			c[y][x].A = 255
		}
	}

	return m
}

func genPic(m image.Image, fn string) {
	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	png.Encode(f, m)
}

func main() {
	genPic(drawPixels(256, 256), "pic.png")
}
