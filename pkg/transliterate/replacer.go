package transliterate

import (
	"bytes"
	"sync"
	"unicode"

	transliterateLang "github.com/Nigh/transliterate/pkg/transliterate-lang"
)

// TODO change Transliterate API

// TODO add comments to exposed types, functions, vars, etc

type Replacer struct {
	Lang      transliterateLang.LangOverwrite
	Data      map[rune][]string
	Separator string
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

	lastType := 0
	changed := false
	langOverwrite, hasLangOverwrite := replacer.Lang[lang]
	for _, char := range text {
		if char < unicode.MaxASCII {
			if unicode.IsSpace(char) {
				lastType = 0
			} else {
				if lastType == 2 {
					buffer.WriteString(replacer.Separator)
				}
				lastType = 1
			}
			buffer.WriteRune(char)
			continue
		}

		changed = true

		if hasLangOverwrite {
			if value, ok := langOverwrite[char]; ok {
				if lastType != 0 {
					buffer.WriteString(replacer.Separator)
				}
				lastType = 2
				buffer.WriteString(value)
				continue
			}
		}

		bank := char >> 8
		code := char & 0xFF

		if value, ok := replacer.Data[bank]; ok {
			if lastType != 0 {
				buffer.WriteString(replacer.Separator)
			}
			lastType = 2
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
