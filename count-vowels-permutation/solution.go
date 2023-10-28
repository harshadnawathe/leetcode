package countvowelspermutation

const mod = 1000000007

func countVowelPermutation(n int) int {
	prevA, prevE, prevI, prevO, prevU := 1, 1, 1, 1, 1

	for i := 2; i <= n; i++ {
		nextA := (prevE + prevI + prevU) % mod
		nextE := (prevA + prevI) % mod
		nextI := (prevE + prevO) % mod
		nextO := prevI
		nextU := (prevO + prevI) % mod

		prevA = nextA
		prevE = nextE
		prevI = nextI
		prevO = nextO
		prevU = nextU
	}

	return (prevA + prevE + prevI + prevO + prevU) % mod
}
