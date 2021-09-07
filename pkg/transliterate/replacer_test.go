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
		Lang: transliterateLang.Data,
		Data: transliterate_data.Data,
	}
	actual := transliterator.Transliterate(text, "de")

	assert.Equal(t, expected, actual)

}

func TestItShouldTransliterateUkrainianCorrectly(t *testing.T) {
	text := "Київ"
	expected := "Kyiv"
	replacer := Replacer{
		Lang: transliterateLang.Data,
		Data: transliterate_data.Data,
	}
	actual := replacer.Transliterate(text, "uk")

	assert.Equal(t, expected, actual)
}

func TestReplacer_Transliterate_withCustomLang(t *testing.T) {
	text := "КиЇв"
	expected := "KyCUv"

	replacer := Replacer{
		Lang: transliterateLang.Data,
		Data: transliterate_data.Data,
	}
	replacer.Lang.AddLanguageOverride("custom", map[rune]string{
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
		Lang: transliterateLang.Data,
		Data: transliterate_data.Data,
	}
	for text, expected := range cases {
		actual := replacer.Transliterate(text, "")
		assert.Equal(t, expected, actual)
	}
}

func BenchmarkReplacer_Transliterate(b *testing.B) {
	text := "Москва выглядит красиво зимой, да?"
	replacer := Replacer{
		Lang: transliterateLang.Data,
		Data: transliterate_data.Data,
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
			Lang: transliterateLang.Data,
			Data: transliterate_data.Data,
		}
	}
}

const theEgg = `
The Egg
By: Andy Weir
Translation: Mariya Solodilova
 
 
Ты умер по пути домой.
Попал в автомобильную аварию. Не особо примечательную, но всё же смертельную. Ты оставил жену и двух детей. Смерть была безболезненная. Скорая пыталась тебя спасти, но всё попусту. Твое тело было так изуродовано, что тебе лучше было уйти, поверь мне.
 
И тогда ты встретил меня.
-- Что… Что произошло?- спросил ты.- Где я?
-- Ты умер, - ответил я, как ни в чем не бывало. Не время жеманничать.
-- Там был… грузовик, и его заносило…
-- Ага,- сказал я.
-- Я… я умер?
-- Ага. Но не расстраивайся, все умирают,- подтвердил я.
 
Ты осмотрелся. Вокруг была пустота. Только ты и я.
-- Что это за место?- спросил ты. – Это жизнь после смерти?
-- Более или менее, - ответил я.
-- А ты бог?
-- Ага,- сказал я. – Я Бог.
-- Моя жена… и дети – пробормотал ты.
-- Что?
-- С ними все нормально?
-- Мне это нравится, - сказал я. – Ты только что погиб и так волнуешься о своей семье. Это очень хорошо.
 
Ты посмотрел на меня с благоговением. В твоих глазах я вовсе не выглядел как Бог. Я казался тебе обычным мужчиной. Или, может быть, женщиной. Каким-то влиятельным человеком с размытым лицом. Скорее учителем начальных классов, чем Господом Всемогущим.
 
-- Не волнуйся, - сказал я. – Они в порядке. Твои дети всегда будут помнить о тебе только лучшее. Они не накопили к тебе неуважение. Твоя жена будет плакать, но в душе будет чувствовать облегчение. Честно говоря, твой брак разваливался. Если тебя это утешит, то могу сказать, что жена твоя будет чувствовать себя очень виноватой за это тайное чувство облегчения.
-- Ооо…- протянул ты. – Ну а что теперь? Ты пошлешь меня в рай или в ад, или что-то вроде того?
-- Ни то, ни другое – ответил я. – Твоя душа переселится в иное тело.
-- Ааа, значит, Индуисты были правы….
-- Все религии правы по-своему – сказал я. – Пойдем со мной.
 
И ты пошел рядом со мной сквозь пустоту.
-- Куда мы идем?
-- Конкретно - никуда. Просто приятно гулять во время разговора.
-- Тогда в чем смысл? – спросил ты. – Когда я буду рожден вновь, я же буду вновь пустым, как стеклышко? Всего лишь дитя. Значит, весь мой опыт и все, чего я добился в той жизни, не будет иметь значения.
-- Вовсе нет! – заверил я. – У тебя внутри уже заложены опыт и мудрость прошлых твоих жизней. Ты просто их в данный момент не помнишь.
 
Я остановился и обнял тебя за плечи.
-- Твоя душа намного огромней, изумительней и прекрасней, чем ты можешь себе представить. Человеческое сознание может воспринимать лишь крошечную долю того, что на самом деле существует. Это словно окунуть палец в стакан воды, чтобы проверить, холодная она или горячая. Ты впускаешь часть себя в этот мир, а когда выходишь из него, то весь  накопленный опыт и знания остаются у тебя.
Ты был в человеке все предыдущие 48 лет, поэтому ты еще не чувствуешь оставшуюся часть своего огромного сознания. Если бы мы с тобой еще здесь походили, ты бы начал постепенно вспоминать все, что было с тобой в прошлых жизнях. Но нет смысла это делать между жизнями.
 
--  Сколько же раз я пережил реинкарнацию?
-- О, много. Очень, очень много. Ты пережил множество разных жизней, - ответил я. – На этот раз ты будешь китайской крестьянкой в 540 году до нашей эры.
-- Подожди, как так? – поперхнулся ты. – Ты посылаешь меня назад во времени?
-- Ну, можно сказать и так. Время в той форме, в которой ты его знаешь, существует только в твоей вселенной. Там, откуда я родом, все происходит по-другому.
-- Откуда ты родом?.. – удивился ты.
-- Ну да, - объяснил я. – Я тоже откуда-то родом. Но совершенно из другого измерения. И там есть еще такие же, как я. Ты, конечно, хочешь знать, каково это там, но, честно говоря, ты  не поймешь.
-- Ааа, -- разочарованно протянул ты. – Но послушай, если я перевоплощаюсь в людей из разного времени, я, наверное, когда-нибудь могу пересечься с самим собой?..
-- Конечно. Такое очень часто происходит. Из-за того, что каждая жизнь осознает лишь себя, ты даже не понимаешь, что встреча произошла.
-- Тогда в чем смысл всего того?
-- Ты серьезно? – удивился я. – Ты спрашиваешь меня, в чем смысл жизни? Немного клише, тебе не кажется?
-- Но это закономерный вопрос, - настойчиво сказал ты.
 
 Я посмотрел тебе в глаза.
-- Смысл жизни, то, ради чего я создал эту вселенную, это чтобы ты развивался.
-- Имеешь в виду человечество? Ты хочешь, чтобы человечество развивалось?
-- Нет-нет, только ты. Я создал всю эту вселенную для тебя. С каждой новой жизнью ты растешь и развиваешься, превращаешься во всеобъемлющий интеллект.
-- Только я? А как же остальные?
-- Остальных не существует. В этой вселенной больше никого нет. Есть только ты и я.
 
Ты уставился на меня.
-- Но все люди на Земле…
-- Это все ты. Разные перевоплощения тебя.
-- Я… Я – ВСЕ?
-- Именно, - с удовлетворением заключил я и похлопал тебя по спине.
-- Я - каждый человек, который когда-либо жил на Земле?
-- И который когда-либо будет жить, да.
-- Я Авраам Линкольн? – поразился ты.
-- И ты Джон Вилкс Бут.
-- Я Гитлер?
-- И ты миллионы его жертв.
-- Я Иисус?
-- И ты каждый из его последователей.
 
Ты замолчал.
-- Каждый раз, причиняя кому-то боль, ты причинял боль самому себе. Каждый раз, делая кому-то добро, ты делал добро себе. Каждый счастливый или грустный момент пребывания на Земле был испытан, или будет испытан только тобой.
 
Ты задумался.
-- Зачем? – наконец спросил ты. – Для чего все это?
-- Потому что однажды ты станешь таким, как я. Потому что ты и есть я. Ты часть меня. Ты дитя моё.
-- Значит, я  и есть Бог? – недоверчиво спросил ты.
-- Нет, пока еще нет. Сейчас ты только зародыш. Ты растешь. Когда ты проживешь каждую человеческую жизнь на Земле во все времена, ты будешь готов родиться.
-- Значит, вся вселенная, - изумленно
сказал ты,- это всего лишь…
-- Яйцо, - подтвердил я. – А теперь тебе пора в новую жизнь.
 
И я отправил тебя в путь.
`

func Benchmark_Transliterate(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for iter := 0; iter < b.N; iter++ {
		_ = Transliterate(theEgg, "")
	}
}

func TestReplacer_Transliterate_okMemory(t *testing.T) {
	text1 := "Москва выглядит красиво зимой, да?"
	text2 := "Мы были активированы мистер Даллиард"
	transliterator := Replacer{
		Lang: transliterateLang.Data,
		Data: transliterate_data.Data,
	}

	out1 := transliterator.Transliterate(text1, "en")
	out2 := transliterator.Transliterate(text2, "en")

	require.Equal(t, "Moskva vygliadit krasivo zimoi, da?", out1)
	require.Equal(t, "My byli aktivirovany mister Dalliard", out2)
	require.NotEqual(t, out1, out2)
}
