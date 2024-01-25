package longestcommonsubsequence

// time: O(m*n)
// space: O(m*n)
func longestCommonSubsequence(text1 string, text2 string) int {
	lcsLength := make([][]int, len(text1)+1)
	for i := range lcsLength {
		lcsLength[i] = make([]int, len(text2)+1)
	}

	for i := len(text1) - 1; i >= 0; i-- {
		for j := len(text2) - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				lcsLength[i][j] = 1 + lcsLength[i+1][j+1]
			} else {
				lcsLength[i][j] = maxOf(lcsLength[i+1][j], lcsLength[i][j+1])
			}
		}
	}
	return lcsLength[0][0]
}

// func longestCommonSubsequence(text1 string, text2 string) int {
// 	var lcsLength func(int, int) int
// 	lcsLength = func(i, j int) int {
// 		if i == len(text1) || j == len(text2) {
//       return 0
//     }
//     if text1[i] == text2[j] {
//       return 1 + lcsLength(i+1, j+1)
//     }
//     return maxOf(lcsLength(i+1, j), lcsLength(i, j+1))
// 	}
//
// 	return lcsLength(0, 0)
// }

func maxOf(lhs, rhs int) int {
	if lhs > rhs {
		return lhs
	}
	return rhs
}
