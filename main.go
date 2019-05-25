package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"image"
	"io/ioutil"
	imageBob "retarded-bob-generator/image"
	"retarded-bob-generator/text"
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

	translated = text.ToBobRetardedString(toTranslate)
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
	font, err := imageBob.LoadFontFromBytes(fontRaw, 20)
	if err != nil {
		return err
	}
	buffer := imageBob.GenerateBobMeme(bobImage, font, translated)
	exportImage(outputPath, buffer.Bytes())
	return nil
}

func exportImage(outputPath string, imageBuffer []byte) {
	ioutil.WriteFile(outputPath, imageBuffer, 0644)
}
