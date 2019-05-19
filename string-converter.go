package main

import (
	"strings"
)

func toBobRetardedString(cleverSentence string) (retardedSentence string) {
	characters := strings.Split(cleverSentence, "")
	for index, char := range characters {
		if index%2 == 0 {
			retardedSentence += strings.ToUpper(char)
		} else {
			retardedSentence += char
		}
	}
	return
}
