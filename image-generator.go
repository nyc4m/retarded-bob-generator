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
	err = board.LoadFontFace("res/font/impact.ttf", 30)
	if err != nil {
		return err
	}
	widthOfText, _ := board.MeasureString(text)
	startWidthOfImage := board.Width()/2 - int(widthOfText/2)
	board.DrawString(text, float64(startWidthOfImage), 350)
	gg.SavePNG(outputPath, board.Image())
	return nil
}
