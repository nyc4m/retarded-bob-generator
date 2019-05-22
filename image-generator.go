package main

import (
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image"
)

func generateImage(bobImage image.Image, fontFace font.Face, text, outputPath string) error {
	board := gg.NewContextForImage(bobImage)
	board.SetHexColor("#ffffff")
	board.SetFontFace(fontFace)
	widthOfText, _ := board.MeasureString(text)
	startWidthOfImage := board.Width()/2 - int(widthOfText/2)
	board.DrawString(text, float64(startWidthOfImage), 350)
	gg.SavePNG(outputPath, board.Image())
	return nil
}
