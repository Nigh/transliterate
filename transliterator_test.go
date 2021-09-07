package transliterator

import (
	"github.com/alexsergivan/transliterator/data"
	"github.com/alexsergivan/transliterator/languages"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItShouldTransliterateGermanCorrectly(t *testing.T) {
	text := "München"
	expected := "Muenchen"

	transliterator := Transliterator{
		LanguageOverrides: languages.LanguageOverridesData,
		Data:              data.TransliterationData,
	}
	actual := transliterator.Transliterate(text, "de")

	assert.Equal(t, expected, actual)

}

func TestItShouldTransliterateUkrainianCorrectly(t *testing.T) {
	text := "Київ"
	expected := "Kyiv"
	transliterator := Transliterator{
		LanguageOverrides: languages.LanguageOverridesData,
		Data:              data.TransliterationData,
	}
	actual := transliterator.Transliterate(text, "uk")

	assert.Equal(t, expected, actual)
}

func TestItShouldTransliterateCorrectlyWithCustomOverrides(t *testing.T) {
	text := "КиЇв"
	expected := "KyCUv"

	transliterator := Transliterator{
		LanguageOverrides: languages.LanguageOverridesData,
		Data:              data.TransliterationData,
	}
	transliterator.LanguageOverrides.AddLanguageOverride("custom", map[rune]string{
		0x407: "CU",
		0x438: "y",
	})
	actual := transliterator.Transliterate(text, "custom")

	assert.Equal(t, expected, actual)
}

func TestItShouldTransliterateGeneral(t *testing.T) {
	cases := map[string]string{
		"北京":           "Bei Jing ",
		"80 km/h":      "80 km/h",
		"дом":          "dom",
		"ⓐⒶ⑳⒇⒛⓴⓾⓿":     "aA20(20)20.20100",
		"ch\u00e2teau": "chateau",
		"\u1eff":       "",
	}

	transliterator := Transliterator{
		LanguageOverrides: languages.LanguageOverridesData,
		Data:              data.TransliterationData,
	}
	for text, expected := range cases {
		actual := transliterator.Transliterate(text, "")
		assert.Equal(t, expected, actual)
	}
}

func BenchmarkTransliterator_Transliterate(b *testing.B) {
	text := "Москва выглядит красиво зимой, да?"
	transliterator := Transliterator{
		LanguageOverrides: languages.LanguageOverridesData,
		Data:              data.TransliterationData,
	}
	var actual string

	b.ReportAllocs()
	b.ResetTimer()

	for iter := 0; iter < b.N; iter++ {
		actual = transliterator.Transliterate(text, "en")
	}

	_ = actual
}

func BenchmarkNewTransliterator(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for iter := 0; iter < b.N; iter++ {
		_ = Transliterator{
			LanguageOverrides: languages.LanguageOverridesData,
			Data:              data.TransliterationData,
		}
	}
}

func TestTransliterator_Transliterate_okMemory(t *testing.T) {
	text1 := "Москва выглядит красиво зимой, да?"
	text2 := "Мы были активированы мистер Даллиард"
	transliterator := Transliterator{
		LanguageOverrides: languages.LanguageOverridesData,
		Data:              data.TransliterationData,
	}

	out1 := transliterator.Transliterate(text1, "en")
	out2 := transliterator.Transliterate(text2, "en")

	require.Equal(t, "Moskva vygliadit krasivo zimoi, da?", out1)
	require.Equal(t, "My byli aktivirovany mister Dalliard", out2)
	require.NotEqual(t, out1, out2)
}
