package image

import (
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
)

//LoadFontFromFile loads the font from a specified path
func LoadFontFromBytes(fontBytes []byte, size float64) (font.Face, error) {
	parsedFont, err := truetype.Parse(fontBytes)

	if err != nil {
		return nil, err
	}

	return truetype.NewFace(parsedFont, &truetype.Options{Size: size}), nil

}
