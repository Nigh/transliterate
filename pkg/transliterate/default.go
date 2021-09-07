package transliterate

var defaultReplacer = Replacer{
	LanguageOverrides: nil,
	Data:              nil,
}

func Transliterate(text, lang string) string {
	return defaultReplacer.Transliterate(text, lang)
}
