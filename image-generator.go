package main

import (
	"github.com/fogleman/gg"
	"image"
)

func generateImage(bobImage image.Image, text, outputPath string) error {
	board := gg.NewContextForImage(bobImage)
	board.SetHexColor("#ffffff")
	err := board.LoadFontFace("res/font/impact.ttf", 30)
	if err != nil {
		return err
	}
	widthOfText, _ := board.MeasureString(text)
	startWidthOfImage := board.Width()/2 - int(widthOfText/2)
	board.DrawString(text, float64(startWidthOfImage), 350)
	gg.SavePNG(outputPath, board.Image())
	return nil
}
