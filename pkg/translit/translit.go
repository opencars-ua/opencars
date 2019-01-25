package translit

func RuneToUA(x rune) (res rune) {
	switch x {
	case 'A':
		res = 'А'
	case 'B':
		res = 'В'
	case 'C':
		res = 'С'
	case 'E':
		res = 'Е'
	case 'H':
		res = 'Н'
	case 'I':
		res = 'І'
	case 'K':
		res = 'К'
	case 'M':
		res = 'М'
	case 'O':
		res = 'О'
	case 'P':
		res = 'Р'
	case 'T':
		res = 'Т'
	case 'X':
		res = 'Х'
	default:
		res = x
	}

	return
}

func ToUA(lexeme string) string {
	runes := make([]rune, 0)

	for _, v := range []rune(lexeme) {
		runes = append(runes, RuneToUA(v))
	}

	return string(runes)
}
