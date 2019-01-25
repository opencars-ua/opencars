package translit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRuneToUA(t *testing.T) {

}

func TestToUA(t *testing.T) {
	t.Run("Test usual plate numbers", func(t *testing.T) {
		assert.Equal(t, "АВС", ToUA("ABC"))
		assert.Equal(t, "123456789", ToUA("123456789"))
		assert.Equal(t, "АА0000АА", ToUA("AA0000AA"))
	})

	t.Run("Test without latin", func(t *testing.T) {
		assert.Equal(t, "АВС", ToUA("АВС"))
		assert.Equal(t, "123456789", ToUA("123456789"))
		assert.Equal(t, "АА0000АА", ToUA("АА0000АА"))
	})

	t.Run("Latin to cyrillic for each region", func(t *testing.T) {
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
