package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

var transNumLit = map[string]string{
	"a": "1",
	"e": "2",
	"i": "3",
	"u": "4",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	textTransliteration := NewTextTransliteration(transNumLit, text)

	fmt.Print(fmt.Sprintf("Case-sensitive Encoding says: %s", textTransliteration.Encode()))
	fmt.Print(fmt.Sprintf("Case-insensitive Encoding says: %s", textTransliteration.EncodeIgnoringCase()))
	fmt.Print(fmt.Sprintf("Case-sensitive Decoding says: %s", textTransliteration.Decode()))
	fmt.Print(fmt.Sprintf("Case-insensitive Decoding says: %s", textTransliteration.DecodeIgnoringCase()))
}

type textTransliteration struct {
	Mapping map[rune]rune
	Text    string
}

func NewTextTransliteration(mapping map[string]string, text string) *textTransliteration {
	return &textTransliteration{Mapping: ToMapOfRunes(mapping), Text: text}
}

type DecodingInterface interface {
	Decode() string
	DecodeIgnoringCase() string
}

type EncodingInterface interface {
	Encode() string
	EncodeIgnoringCase() string
}

func (t *textTransliteration) Encode() string {
	return string(ReplaceRunesAsPerMapping([]rune(t.Text), t.Mapping, true))
}

func (t *textTransliteration) EncodeIgnoringCase() string {
	return string(ReplaceRunesAsPerMapping([]rune(t.Text), t.Mapping, false))
}

func (t *textTransliteration) DecodeIgnoringCase() string {
	return string(ReplaceRunesAsPerMapping([]rune(t.Text), t.reverseMapping(), true))
}

func (t *textTransliteration) Decode() string {
	return string(ReplaceRunesAsPerMapping([]rune(t.Text), t.reverseMapping(), false))
}

func ReplaceRunesAsPerMapping(runes []rune, m map[rune]rune, caseSensitive bool) []rune {

	for i, r := range runes {
		if v, ok := m[r]; ok {
			runes[i] = v
			continue
		}

		if !caseSensitive {
			if unicode.IsUpper(r) {
				if v, ok := m[unicode.ToLower(r)]; ok {
					runes[i] = v
				}
			} else {
				if v, ok := m[unicode.ToUpper(r)]; ok {
					runes[i] = v
				}
			}
		}
	}

	return runes
}

func (t *textTransliteration) reverseMapping() map[rune]rune {
	return Reverse(t.Mapping)
}

func Reverse(m map[rune]rune) map[rune]rune {
	n := make(map[rune]rune)

	for k, v := range m {
		n[v] = k
	}

	return n
}

func ToMapOfRunes(m map[string]string) map[rune]rune {
	n := make(map[rune]rune)

	for k, v := range m {
		kRunes := []rune(k)
		if len(kRunes) > 1 {
			panic("One char keys only")
		}

		vRunes := []rune(v)
		if len(vRunes) > 1 {
			panic("One char values only")
		}

		n[kRunes[0]] = vRunes[0]
	}

	return n
}
