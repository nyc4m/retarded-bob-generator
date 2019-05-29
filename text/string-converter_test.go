package text

import (
	"testing"
)

func TestToBobRetardedString(t *testing.T) {
	test := []struct {
		input  string
		output string
	}{
		{
			input:  "salut",
			output: "SaLuT",
		},
		{
			input:  "Je tenais à partager",
			output: "Je TeNaIs À pArTaGeR",
		},
	}
	for _, element := range test {
		res := ToBobRetardedString(element.input)
		if res != element.output {
			t.Log(res + " is different than " + element.output)
			t.Fail()
		}
	}
}
