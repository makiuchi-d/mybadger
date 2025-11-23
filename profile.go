package main

import (
	"errors"

	"tinygo.org/x/tinydraw"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

func BadgeProfile(company, handle, name, prof1, prof2 string, img []byte) func() error {
	return func() error {
		err := DrawFrame(img, Title{{&freesans.Bold9pt7b, company}})
		if err != nil {
			return err
		}

		tinydraw.Line(&display, 0, 87, WIDTH-IMGW, 87, black)
		tinydraw.Line(&display, 0, 107, WIDTH-IMGW, 107, black)

		tinyfont.WriteLine(&display, &freesans.Bold12pt7b, 8, 56, handle, black)
		tinyfont.WriteLine(&display, &freesans.Bold9pt7b, 8, 78, name, black)

		_, ow := tinyfont.LineWidth(&freesans.Regular9pt7b, prof1)
		x := int16(WIDTH-IMGW-ow) / 2
		if x < 0 {
			return errors.New("prof1 too long")
		}
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, x, 102, prof1, black)

		_, ow = tinyfont.LineWidth(&freesans.Regular9pt7b, prof2)
		x = int16(WIDTH-IMGW-ow) / 2
		if x < 0 {
			return errors.New("prof2 too long")
		}
		tinyfont.WriteLine(&display, &freesans.Regular9pt7b, x, 122, prof2, black)

		return nil
	}
}
