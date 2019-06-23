package text

import (
	"strings"
)

//ToBobRetardedString Translate to a bob retarded string
//so 'Hello world', becomes 'HeLlO wOrLd'
func ToBobRetardedString(cleverSentence string) (retardedSentence string) {
	words := strings.Split(cleverSentence, " ")
	sentence := make([]string, 0)
	startWithUpper := true
	for _, word := range words {
		var retardedWord string
		retardedWord, startWithUpper = toBobRetardedWord(word, startWithUpper)
		sentence = append(sentence, retardedWord)
	}
	return strings.Join(sentence, " ")
}

func toBobRetardedWord(cleverWord string, beginWithUpper bool) (retardedSentence string, shouldStartWithUpper bool) {
	characters := strings.Split(strings.ToLower(cleverWord), "")
	if beginWithUpper == false {
		retardedSentence = characters[0]
		characters = characters[1:]
	}

	for index, char := range characters {
		if index%2 == 0 {
			retardedSentence += strings.ToUpper(char)
			shouldStartWithUpper = false
		} else {
			retardedSentence += char
			shouldStartWithUpper = true
		}
	}
	return
}
