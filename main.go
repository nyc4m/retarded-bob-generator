package main

import (
	"flag"
	"fmt"
)

func main() {

	var outputPath, toTranslate string
	var toText bool
	var translated string

	flag.StringVar(&toTranslate, "input", "hello world", "string that should be convert to bob retarded language")
	flag.StringVar(&outputPath, "out", "./bob", "string indicating where the file image should be generated")
	flag.BoolVar(&toText, "text", true, "if true, will only ouptut to text")

	flag.Parse()

	translated = toBobRetardedString(toTranslate)
	switch toText {
	case true:
		fmt.Println(translated)
	case false:
		generateImage("./bob.png", translated, outputPath)
	}
}
