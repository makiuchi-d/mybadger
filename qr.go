package main

import (
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"tinygo.org/x/tinyfont"
)

func BadgeQR(msg string, title Title, body Body) func() error {
	bmx, err := genQR(msg)
	if err != nil {
		return func() error { return err }
	}
	bits, err := bmxToBits(bmx)
	if err != nil {
		return func() error { return err }
	}
	return func() error {
		if err := DrawFrame(bits, title); err != nil {
			return err
		}

		for _, b := range body {
			tinyfont.WriteLine(&display, b.Font, b.X, b.Y, b.Str, black)
		}

		return nil
	}
}

func genQR(msg string) (*gozxing.BitMatrix, error) {
	hints := map[gozxing.EncodeHintType]any{
		gozxing.EncodeHintType_ERROR_CORRECTION: "M",
	}

	bmx, err := qrcode.NewQRCodeWriter().Encode(msg, gozxing.BarcodeFormat_QR_CODE, IMGW, IMGH, hints)
	if err != nil {
		return nil, err
	}

	return bmx, nil
}

func bmxToBits(bmx *gozxing.BitMatrix) ([]byte, error) {
	w, h := bmx.GetWidth(), bmx.GetHeight()
	bits := make([]byte, w*h/8)
	count := 0
	for y := range h {
		for x := range w {
			if bmx.Get(x, y) {
				bits[count/8] |= (1 << (7 - count%8))
			}
			count++
		}
	}
	return bits, nil
}
