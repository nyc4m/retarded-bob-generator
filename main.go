package main

import (
	"flag"
	"fmt"
	"github.com/fogleman/gg"
)

func main() {

	var outputPath, toTranslate string
	var toText bool
	var translated string

	flag.StringVar(&toTranslate, "input", "hello world", "string that should be convert to bob retarded language")
	flag.StringVar(&outputPath, "out", "./bob.png", "string indicating where the file image should be generated")
	flag.BoolVar(&toText, "text", false, "if true, will only ouptut to text")

	flag.Parse()

	translated = toBobRetardedString(toTranslate)
	switch toText {
	case true:
		fmt.Println(translated)
	case false:
		bobImage, err := gg.LoadPNG("./res/bob_source.png")
		if err != nil {
			fmt.Println(err.Error())
		}
		generateImage(bobImage, translated, outputPath)
	}
}
