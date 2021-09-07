package transliterate_lang

type LangOverwrite map[string]map[rune]string

var Data = LangOverwrite{
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
