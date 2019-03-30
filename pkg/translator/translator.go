package translator

// RuneToUA returns Ukrainian equivalent of Latin character.
func RuneToUA(char rune) rune {
	switch char {
	case 'A':
		return 'А'
	case 'B':
		return 'В'
	case 'C':
		return 'С'
	case 'E':
		return 'Е'
	case 'H':
		return 'Н'
	case 'I':
		return 'І'
	case 'K':
		return 'К'
	case 'M':
		return 'М'
	case 'O':
		return 'О'
	case 'P':
		return 'Р'
	case 'T':
		return 'Т'
	case 'X':
		return 'Х'
	default:
		return char
	}
}

// ToUA translates string in Ukrainian to Latin equivalent.
func ToUA(lexeme string) string {
	chars := make([]rune, 0)

	for _, v := range lexeme {
		chars = append(chars, RuneToUA(v))
	}

	return string(chars)
}
