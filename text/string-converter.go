package text

import (
	"strings"
)

//ToBobRetardedString Translate to a bob retarded string
//so 'Hello world', becomes 'HeLlO wOrLd'
func ToBobRetardedString(cleverSentence string) (retardedSentence string) {
	characters := strings.Split(strings.ToLower(cleverSentence), "")
	for index, char := range characters {
		if index%2 == 0 {
			retardedSentence += strings.ToUpper(char)
		} else {
			retardedSentence += char
		}
	}
	return
}
