package main

import (
	"image/color"

	"tinygo.org/x/drivers/pixel"
	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
)

const (
	WIDTH  = 296
	HEIGHT = 128

	IMGW = 128
	IMGH = 128
	IMGX = WIDTH - IMGW
	IMGY = 0

	TITLEH = 30
)

var (
	black = color.RGBA{1, 1, 1, 255}
	white = color.RGBA{0, 0, 0, 255}
)

type TitleElement struct {
	Font tinyfont.Fonter
	Str  string
}

type Title []TitleElement

type BodyElement struct {
	X    int16
	Y    int16
	Font tinyfont.Fonter
	Str  string
}

type Body []BodyElement

func DrawFrame(img []byte, title Title) error {
	// 先に画像、枠をあとから描く
	err := display.DrawBitmap(IMGX, IMGY, pixel.NewImageFromBytes[pixel.Monochrome](IMGW, IMGH, img))
	if err != nil {
		return err
	}

	DrawLine(0, 0, WIDTH-1, 0)
	DrawLine(0, 0, 0, HEIGHT-1)
	DrawLine(WIDTH-1, 0, WIDTH-1, HEIGHT-1)
	DrawLine(0, HEIGHT-1, WIDTH-1, HEIGHT-1)
	DrawLine(IMGX, IMGY, IMGX, HEIGHT-1)

	fillRect(0, 0, IMGX, TITLEH)
	drawTitle(title)

	return nil
}

func DrawLine(x0, y0, x1, y1 int16) {
	tinydraw.Line(&display, x0, y0, x1, y1, black)
}

func fillRect(x int16, y int16, w int16, h int16) {
	for j := y; j < y+h; j++ {
		for i := x; i < x+w; i++ {
			display.SetPixel(i, j, black)
		}
	}
}

func drawTitle(title Title) {
	const y = TITLEH - 8
	x := int16(8)
	for _, t := range title {
		tinyfont.WriteLine(&display, t.Font, x, y, t.Str, white)
		_, w := tinyfont.LineWidth(t.Font, t.Str)
		x += int16(w) + 4
	}
}
