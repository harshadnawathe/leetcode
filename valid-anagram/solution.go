package validanagram

// // time: O(n), space:O(1)
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
// 		return false
// 	}
//
// 	charCntS, charCntT := [26]int{}, [26]int{}
// 	for i := 0; i < len(s); i++ {
// 		charCntS[s[i]-'a'] += 1
// 		charCntT[t[i]-'a'] += 1
// 	}
//
// 	for i := 0; i < 26; i++ {
// 		if charCntS[i] != charCntT[i] {
// 			return false
// 		}
// 	}
//
// 	return true
// }

// // time: O(n), space:O(1)
// func isAnagram(s string, t string) bool {
// 	ccs := charCount(s)
// 	cct := charCount(t)
//
// 	return reflect.DeepEqual(ccs, cct)
// }
//
// func charCount(s string) []int {
// 	cc := make([]int, 26)
// 	for i := 0; i < len(s); i++ {
// 		cc[s[i]-'a'] += 1
// 	}
// 	return cc
// }

// time: O(n), space: O(1)
func isAnagram(s string, t string) bool {
	cc := make([]int, 26)

	for i := 0; i < len(s); i++ {
		cc[s[i]-'a'] += 1
	}

	for i := 0; i < len(t); i++ {
		cc[t[i]-'a'] -= 1
	}

	for _, n := range cc {
		if n != 0 {
			return false
		}
	}
	return true
}
