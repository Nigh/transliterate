package languages

// LanguageOverrides structure.
type LanguageOverrides map[string]map[rune]string

// AddLanguageOverride adds custom transliteration overrides for specific language.
func (lo LanguageOverrides) AddLanguageOverride(langcode string, override map[rune]string) {
	lo[langcode] = override
}

var LanguageOverridesData = LanguageOverrides{
	"de": DE,
	"da": DA,
	"eo": EO,
	"ru": RU,
	"bg": BG,
	"sv": SV,
	"hu": HU,
	"hr": HR,
	"sl": SL,
	"sr": SR,
	"nb": NB,
	"uk": UK,
	"mk": MK,
	"ca": CA,
	"bs": BS,
}
