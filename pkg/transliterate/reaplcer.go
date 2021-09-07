package transliterate

import (
	"bytes"
	"unicode"

	transliterateLang "github.com/alexsergivan/transliterator/pkg/transliterate-lang"
)

// Replacer structure.
type Replacer struct {
	LanguageOverrides transliterateLang.LanguageOverrides
	Data              map[rune][]string
}

// Transliterate performs transliteration of the input text. If the lang (ISO 639-1) is specified, it will use
// specific language transliteration rules.
func (replacer *Replacer) Transliterate(text, lang string) string {
	memory := memoryPool.Get().(*[]byte)
	defer memoryPool.Put(memory)
	buffer := bytes.NewBuffer(*memory)
	buffer.Reset()

	for _, char := range text {
		if overrides, ok := replacer.LanguageOverrides[lang]; ok {
			if val, ok := overrides[char]; ok {
				buffer.WriteString(val)
				continue
			}
		}

		// If the char number less then maximum ASCII value, use it directly.
		if char < unicode.MaxASCII {
			buffer.WriteRune(char)
			continue
		}

		// Example: "Ї" => in the hexadecimal - 0x407
		// bank: 0x4
		// code: 0x7

		// Shifting char to the right by 8 bits.
		bank := char >> 8

		// masks the variable so it leaves only the value in the last 8 bits, and ignores all the rest of the bits
		code := char & 0xff
		if transliterationDataVal, ok := replacer.Data[bank]; ok {
			if len(transliterationDataVal) > int(code) {
				buffer.WriteString(transliterationDataVal[code])
			}
		}
	}

	return buffer.String()
}
