package translator

import "testing"

func TestRuneToUA(t *testing.T) {
	if 'А' != RuneToUA('A') {
		t.Fail()
	}

	if 'В' != RuneToUA('B') {
		t.Fail()
	}

	if 'С' != RuneToUA('C') {
		t.Fail()
	}

	if 'Z' != RuneToUA('Z') {
		t.Fail()
	}

	if '1' != RuneToUA('1') {
		t.Fail()
	}
}

func TestToUA(t *testing.T) {
	t.Run("simple strings", func(t *testing.T) {
		if "АВС" != ToUA("ABC") {
			t.Fail()
		}

		if "АА0000АА" != ToUA("AA0000AA") {
			t.Fail()
		}
	})

	t.Run("without latin", func(t *testing.T) {
		if "АБВ" != ToUA("АБВ") {
			t.Fail()
		}

		if "123456789" != ToUA("123456789") {
			t.Fail()
		}

		if "АХ1234ВА" != ToUA("АХ1234ВА") {
			t.Fail()
		}
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
			done := ToUA(fixtures[i])

			if expected[i] != ToUA(fixtures[i]) {
				t.Errorf("%s not equal %s", expected, done)
			}
		}
	})
}

func BenchmarkToUA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ToUA("AA1728BP")
	}
}

func BenchmarkRuneToUA(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RuneToUA('C')
	}
}