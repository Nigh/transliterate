package transliterate

import (
	"bytes"
	"sync"
	"unicode"

	transliterateLang "github.com/alexsergivan/transliterator/pkg/transliterate-lang"
)

type Replacer struct {
	Lang transliterateLang.LangOverwrite
	Data map[rune][]string
}

var bufferPool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(nil)
	},
}

// Transliterate performs transliteration of the input text. If the lang (ISO 639-1) is specified, it will use specific
// language transliteration rules.
func (replacer *Replacer) Transliterate(text, lang string) string {
	buffer := bufferPool.Get().(*bytes.Buffer)
	defer bufferPool.Put(buffer)
	buffer.Reset()

	changed := false
	langOverwrite, hasLangOverwrite := replacer.Lang[lang]
	for _, char := range text {
		if char < unicode.MaxASCII {
			buffer.WriteRune(char)
			continue
		}

		changed = true

		if hasLangOverwrite {
			if value, ok := langOverwrite[char]; ok {
				buffer.WriteString(value)
				continue
			}
		}

		bank := char >> 8
		code := char & 0xFF

		if value, ok := replacer.Data[bank]; ok {
			if len(value) > int(code) {
				buffer.WriteString(value[code])
			}
		}
	}

	if !changed {
		return text
	}

	return buffer.String()
}
