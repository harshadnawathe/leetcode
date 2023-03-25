package romantointeger

var romanToArabic = map[byte]int{
	'M': 1000,
	'D': 500,
	'C': 100,
	'L': 50,
	'X': 10,
	'V': 5,
	'I': 1,
}

func romanToInt(s string) int {

	num := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romanToArabic[s[i+1]] > romanToArabic[s[i]] {
			num += romanToArabic[s[i+1]] - romanToArabic[s[i]]
			i++
		} else {
			num += romanToArabic[s[i]]
		}
	}
	return num
}

func romanToInt1(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case 'M':
			result += 1000
		case 'D':
			result += 500
		case 'C':
			if i+1 < len(s) && s[i+1] == 'M' {
				result += 900
				i++
			} else if i+1 < len(s) && s[i+1] == 'D' {
				result += 400
				i++
			} else {
				result += 100
			}
		case 'L':
			result += 50
		case 'X':
			if i+1 < len(s) && s[i+1] == 'C' {
				result += 90
				i++
			} else if i+1 < len(s) && s[i+1] == 'L' {
				result += 40
				i++
			} else {
				result += 10
			}
		case 'V':
			result += 5
		case 'I':
			if i+1 < len(s) && s[i+1] == 'X' {
				result += 9
				i++
			} else if i+1 < len(s) && s[i+1] == 'V' {
				result += 4
				i++
			} else {
				result += 1
			}
		}
	}

	return result
}
