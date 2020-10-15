package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"os"
)

const size = 240

var (
	white = color.Gray{Y: 0xff}
	black = color.Gray{Y: 0}
)

func whiten(img draw.Image) {
	rect := img.Bounds()
	for x := rect.Min.X; x < rect.Max.X; x++ {
		for y := rect.Min.Y; y < rect.Max.Y; y++ {
			img.Set(x, y, white)
		}
	}
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func drawLine(img draw.Image, color color.Color, x0, y0, x1, y1 int) {
	dx := abs(x1 - x0)
	dy := abs(y1 - y0)
	var sx, sy int
	if x0 < x1 {
		sx = 1
	} else {
		sx = -1
	}
	if y0 < y1 {
		sy = 1
	} else {
		sy = -1
	}
	err := dx - dy

	var e2 int
	for {
		img.Set(x0, y0, color)
		if x0 == x1 && y0 == y1 {
			return
		}
		e2 = 2 * err
		if e2 > -dy {
			err = err - dy
			x0 = x0 + sx
		}
		if e2 < dx {
			err = err + dx
			y0 = y0 + sy
		}
	}
}

func drawCamp(img draw.Image) {
	rect := img.Bounds()
	w := rect.Max.X - rect.Min.X
	h := rect.Max.Y - rect.Min.Y
	ws := w / 16
	hs := h / 16
	x1 := ws * 4
	x2 := ws * 12
	x3 := ws * 2
	x4 := ws * 6
	x5 := ws * 14
	y1 := hs * 5
	y2 := hs * 11
	drawLine(img, black, x1, y1, x2, y1)
	drawLine(img, black, x1, y1, x3, y2)
	drawLine(img, black, x1, y1, x1, y2)
	drawLine(img, black, x1, y1, x4, y2)
	drawLine(img, black, x2, y1, x5, y2)
	drawLine(img, black, 0, y2, w, y2)
}

func save(avatar image.Image) {
	f, err := os.Create(fmt.Sprintf("avatar-%d.jpeg", size))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = jpeg.Encode(f, avatar, &jpeg.Options{Quality: 100})
	if err != nil {
		panic(err)
	}
}

func main() {
	flag.Parse()

	avatar := image.NewGray(image.Rect(0, 0, size, size))
	whiten(avatar)
	drawCamp(avatar)

	save(avatar)
}
