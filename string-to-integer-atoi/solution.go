package stringtointegeratoi

const (
	uintMin = uint32(0)
	uintMax = ^uintMin
	intMax  = int(uintMax >> 1)
	intMin  = -intMax - 1

	intMaxBy10      = intMax / 10
	intMaxLastDigit = intMax % 10
	intMinBy10      = intMin / 10
	intMinLastDigit = intMin % 10
)

func myAtoi(s string) int {
	value := 0
	sign := 1
	expectedToken := tokenTypeWhitespace | tokenTypeSign | tokenTypeDigit
loop:
	for i := 0; i < len(s); i++ {
		tokType := tokenTypeOf(s[i])

		if tokType&expectedToken == 0 {
			break loop
		}

		switch tokType {
		case tokenTypeWhitespace:
		case tokenTypeSign:
			expectedToken = tokenTypeDigit
			if s[i] == '-' {
				sign = -1
			}
		case tokenTypeDigit:
			expectedToken = tokenTypeDigit
			digit := int(s[i] - '0')
			if sign == -1 {
				num := sign * value
				if num < intMinBy10 || (num == intMinBy10 && digit*sign < intMinLastDigit) {
					sign = 1
					value = intMin
					break loop
				}
			} else {
				if value > intMaxBy10 || (value == intMaxBy10 && digit > intMaxLastDigit) {
					value = intMax
					break loop
				}
			}
			value = value*10 + digit
		default:
			break loop
		}
	}
	return sign * value
}

type tokenType int

const (
	tokenTypeWhitespace tokenType = 1 << iota
	tokenTypeSign
	tokenTypeDigit
	tokenTypeOther
)

func tokenTypeOf(b byte) tokenType {
	if b == ' ' || b == '\t' {
		return tokenTypeWhitespace
	}

	if b == '+' || b == '-' {
		return tokenTypeSign
	}

	if b >= '0' && b <= '9' {
		return tokenTypeDigit
	}

	return tokenTypeOther
}
