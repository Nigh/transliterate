package transliterate

import (
	"bytes"
	"unicode"

	"github.com/alexsergivan/transliterator/internal"

	transliterateLang "github.com/alexsergivan/transliterator/pkg/transliterate-lang"
)

// Replacer structure.
type Replacer struct {
	Lang transliterateLang.LangOverwrite
	Data map[rune][]string
}

// Transliterate performs transliteration of the input text. If the lang (ISO 639-1) is specified, it will use specific
// language transliteration rules.
func (replacer *Replacer) Transliterate(text, lang string) string {
	memory := make([]byte, 0, len(text))
	buffer := bytes.NewBuffer(memory)

	langOverwrite, hasLangOverwrite := replacer.Lang[lang]
	for _, char := range text {
		if hasLangOverwrite {
			if value, ok := langOverwrite[char]; ok {
				buffer.WriteString(value)
				continue
			}
		}

		if char < unicode.MaxASCII {
			buffer.WriteRune(char)
			continue
		}

		bank := char >> 8
		code := char & 0xFF

		if value, ok := replacer.Data[bank]; ok {
			if len(value) > int(code) {
				buffer.WriteString(value[code])
			}
		}
	}

	return internal.BytesToString(memory)
}
