package transliterate

import (
	transliterateData "github.com/Nigh/transliterate/pkg/transliterate-data"
	transliterateLang "github.com/Nigh/transliterate/pkg/transliterate-lang"
)

var defaultReplacer = Replacer{
	Lang: transliterateLang.Data,
	Data: transliterateData.Data,
}

func Sugar(s, lang string) func(in string) string {
	var defaultReplacer = Replacer{
		Lang:      transliterateLang.Data,
		Data:      transliterateData.Data,
		Separator: s,
	}
	return func(text string) string {
		return defaultReplacer.Transliterate(text, lang)
	}
}

// Transliterate is a helper function around a default Replacer using the transliterate_data.Data and transliterate_lang.Data.
func Transliterate(text, lang string) string {
	return defaultReplacer.Transliterate(text, lang)
}
