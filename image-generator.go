package main

import (
	"github.com/fogleman/gg"
)

func generateImage(bobPath, text, outputPath string) error {
	image, err := gg.LoadPNG(bobPath)
	if err != nil {
		return err
	}
	board := gg.NewContextForImage(image)
	board.SetHexColor("#ffffff")
	err = board.LoadFontFace("font/impact.ttf", 30)
	if err != nil {
		return err
	}
	board.DrawString(text, 100, 350)
	gg.SavePNG(outputPath, board.Image())
	return nil
}
