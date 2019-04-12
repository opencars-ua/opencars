package translator

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneToUA(t *testing.T) {
	assert.Equal(t, 'А', RuneToUA('A'))
	assert.Equal(t, 'В', RuneToUA('B'))
	assert.Equal(t, 'С', RuneToUA('C'))
	assert.Equal(t, 'Z', RuneToUA('Z'))
	assert.Equal(t, '1', RuneToUA('1'))
}

func TestToUA(t *testing.T) {
	t.Run("simple strings", func(t *testing.T) {
		assert.Equal(t, "АВС", ToUA("ABC"))
		assert.Equal(t, "АА0000АА", ToUA("AA0000AA"))
	})

	t.Run("without latin", func(t *testing.T) {
		assert.Equal(t, "АБВ", "АБВ")
		assert.Equal(t, "123456789", "123456789")
		assert.Equal(t, "АХ1234ВА", "АХ1234ВА")
	})

	t.Run("latin to cyrillic for each region", func(t *testing.T) {
		fixtures := []string{
			"AK", "AB", "AC", "AE", "AH", "AM", "AO", "AP", "AT",
			"AA", "AI", "BA", "BB", "BC", "BE", "BH", "BI", "BK",
			"CH", "BM", "BO", "AX", "BT", "BX", "CA", "CB", "CE",
		}

		expected := []string{
			"АК", "АВ", "АС", "АЕ", "АН", "АМ", "АО", "АР", "АТ",
			"АА", "АІ", "ВА", "ВВ", "ВС", "ВЕ", "ВН", "ВІ", "ВК",
			"СН", "ВМ", "ВО", "АХ", "ВТ", "ВХ", "СА", "СВ", "СЕ",
		}

		for i := range fixtures {
			assert.Equal(t, expected[i], ToUA(fixtures[i]))
		}
	})
}

func BenchmarkToUA(b *testing.B) {
	fixtures := make([]string, b.N)

	for i := 0; i < b.N; i++ {
		fixtures[i] = strconv.Itoa(1000 + rand.Int()%8999)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		ToUA("AA" + fixtures[i] + "BP")
	}
}

func BenchmarkRuneToUA(b *testing.B) {
	fixtures := make([]rune, b.N)

	for i := 0; i < b.N; i++ {
		fixtures[i] = rune('A' + rand.Int()%'Z')
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		RuneToUA(fixtures[i])
	}
}
