package image

import (
	"bytes"
	"github.com/fogleman/gg"
	"golang.org/x/image/font"
	"image"
	"image/png"
)

//GenerateBobMeme generates the famous meme. it returns a bytes.Buffer
//for facilitating exporting
func GenerateBobMeme(bobImage image.Image, fontFace font.Face, text string) *bytes.Buffer {
	board := gg.NewContextForImage(bobImage)
	board.SetHexColor("#ffffff")
	board.SetFontFace(fontFace)
	widthOfText, _ := board.MeasureString(text)
	startWidthOfImage := board.Width()/2 - int(widthOfText/2)
	board.DrawString(text, float64(startWidthOfImage), 350)

	var buffer bytes.Buffer
	png.Encode(&buffer, board.Image())
	return &buffer
}
