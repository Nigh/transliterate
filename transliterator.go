package transliterator

import (
	"bytes"
	"sync"
	"unicode"

	"github.com/alexsergivan/transliterator/data"
	"github.com/alexsergivan/transliterator/languages"
)

// Transliterator structure.
type Transliterator struct {
	LanguageOverrides *languages.LanguageOverrides
	Data              map[rune][]string
}

// NewTransliterator creates Transliterator object.
func NewTransliterator(customLanguageOverrides *map[string]map[rune]string) *Transliterator {
	languageOverrides := languages.NewLanguageOverrides()
	if customLanguageOverrides != nil {
		languageOverrides.AddLanguageOverrides(customLanguageOverrides)
	}

	return &Transliterator{
		LanguageOverrides: languageOverrides,
		Data:              data.NewTransliterationData().Data,
	}
}

var memoryPool = sync.Pool{
	New: func() interface{} {
		m := make([]byte, 1024)
		return &m
	},
}

// Transliterate performs transliteration of the input text. If the langcode (ISO 639-1) is specified, it will use
// specific language transliteration rules.
func (t *Transliterator) Transliterate(text, langcode string) string {
	memory := memoryPool.Get().(*[]byte)
	defer memoryPool.Put(memory)
	buffer := bytes.NewBuffer(*memory)
	buffer.Reset()

	for _, char := range text {
		if overrides, ok := t.LanguageOverrides.Overrides[langcode]; ok {
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
		if transliterationDataVal, ok := t.Data[bank]; ok {
			if len(transliterationDataVal) > int(code) {
				buffer.WriteString(transliterationDataVal[code])
			}
		}
	}

	return buffer.String()
}
