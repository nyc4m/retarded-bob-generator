package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/golang/freetype/truetype"
	"image"
)

func main() {

	var outputPath, toTranslate string
	var toText bool
	var translated string

	fileBox := rice.MustFindBox("./res")

	flag.StringVar(&toTranslate, "input", "hello world", "string that should be convert to bob retarded language")
	flag.StringVar(&outputPath, "out", "./bob.png", "string indicating where the file image should be generated")
	flag.BoolVar(&toText, "text", false, "if true, will only ouptut to text")

	flag.Parse()

	translated = toBobRetardedString(toTranslate)
	switch toText {
	case true:
		fmt.Println(translated)
	case false:
		err := handleImageGeneration(fileBox, translated, outputPath)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func handleImageGeneration(fileBox *rice.Box, translated, outputPath string) error {
	imageRaw, err := fileBox.Bytes("bob_source.png")
	fontRaw, err := fileBox.Bytes("font/impact.ttf")
	if err != nil {
		return err
	}
	bobImage, _, err := image.Decode(bytes.NewBuffer(imageRaw))
	parsedFont, err := truetype.Parse(fontRaw)
	fontFace := truetype.NewFace(parsedFont, &truetype.Options{Size: 50})
	if err != nil {
		return err
	}
	generateImage(bobImage, fontFace, translated, outputPath)
	return nil
}
