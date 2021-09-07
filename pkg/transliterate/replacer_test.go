package transliterate

import (
	"github.com/alexsergivan/transliterator/pkg/transliterate-data"
	transliterateLang "github.com/alexsergivan/transliterator/pkg/transliterate-lang"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldReplaceGermanCorrectly(t *testing.T) {
	text := "München"
	expected := "Muenchen"

	transliterator := Replacer{
		LanguageOverrides: transliterateLang.LanguageOverridesData,
		Data:              transliterate_data.TransliterationData,
	}
	actual := transliterator.Transliterate(text, "de")

	assert.Equal(t, expected, actual)

}

func TestItShouldTransliterateUkrainianCorrectly(t *testing.T) {
	text := "Київ"
	expected := "Kyiv"
	replacer := Replacer{
		LanguageOverrides: transliterateLang.LanguageOverridesData,
		Data:              transliterate_data.TransliterationData,
	}
	actual := replacer.Transliterate(text, "uk")

	assert.Equal(t, expected, actual)
}

func TestReplacer_Transliterate_withCustomLang(t *testing.T) {
	text := "КиЇв"
	expected := "KyCUv"

	replacer := Replacer{
		LanguageOverrides: transliterateLang.LanguageOverridesData,
		Data:              transliterate_data.TransliterationData,
	}
	replacer.LanguageOverrides.AddLanguageOverride("custom", map[rune]string{
		0x407: "CU",
		0x438: "y",
	})
	actual := replacer.Transliterate(text, "custom")

	assert.Equal(t, expected, actual)
}

func TestReplacer_Transliterate_general(t *testing.T) {
	cases := map[string]string{
		"北京":           "Bei Jing ",
		"80 km/h":      "80 km/h",
		"дом":          "dom",
		"ⓐⒶ⑳⒇⒛⓴⓾⓿":     "aA20(20)20.20100",
		"ch\u00e2teau": "chateau",
		"\u1eff":       "",
	}

	replacer := Replacer{
		LanguageOverrides: transliterateLang.LanguageOverridesData,
		Data:              transliterate_data.TransliterationData,
	}
	for text, expected := range cases {
		actual := replacer.Transliterate(text, "")
		assert.Equal(t, expected, actual)
	}
}

func BenchmarkReplacer_Transliterate(b *testing.B) {
	text := "Москва выглядит красиво зимой, да?"
	replacer := Replacer{
		LanguageOverrides: transliterateLang.LanguageOverridesData,
		Data:              transliterate_data.TransliterationData,
	}
	var actual string

	b.ReportAllocs()
	b.ResetTimer()

	for iter := 0; iter < b.N; iter++ {
		actual = replacer.Transliterate(text, "en")
	}

	_ = actual
}

func BenchmarkNewReplacer(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for iter := 0; iter < b.N; iter++ {
		_ = Replacer{
			LanguageOverrides: transliterateLang.LanguageOverridesData,
			Data:              transliterate_data.TransliterationData,
		}
	}
}

func TestReplacer_Transliterate_okMemory(t *testing.T) {
	text1 := "Москва выглядит красиво зимой, да?"
	text2 := "Мы были активированы мистер Даллиард"
	transliterator := Replacer{
		LanguageOverrides: transliterateLang.LanguageOverridesData,
		Data:              transliterate_data.TransliterationData,
	}

	out1 := transliterator.Transliterate(text1, "en")
	out2 := transliterator.Transliterate(text2, "en")

	require.Equal(t, "Moskva vygliadit krasivo zimoi, da?", out1)
	require.Equal(t, "My byli aktivirovany mister Dalliard", out2)
	require.NotEqual(t, out1, out2)
}
