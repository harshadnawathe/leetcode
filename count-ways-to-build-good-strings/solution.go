package countwaystobuildgoodstrings

const mod = 1000000007

// tabular dp
func countGoodStrings(low int, high int, zero int, one int) int {
	count := make([]int, high+1)
	count[0] = 1

	for n := 1; n <= high; n++ {
		if n >= zero {
			count[n] += count[n-zero]
		}

		if n >= one {
			count[n] += count[n-one]
		}

		count[n] %= mod
	}

	total := 0
	for n := low; n <= high; n++ {
		total += count[n]
		total %= mod
	}

	return total
}

// recursion with memoization
func countGoodStrings2(low int, high int, zero int, one int) int {
	cache := make(map[int]int)
	var f func(int) int
	f = func(n int) int {
		if n == 0 {
			return 1
		}
		if n < 0 {
			return 0
		}
		if cached, found := cache[n]; found {
			return cached
		}

		cache[n] = f(n-zero) + f(n-one)
		cache[n] %= mod
		return cache[n]
	}
	count := 0
	for n := low; n <= high; n++ {
		count += f(n)
		count %= mod
	}
	return count
}

// simple recursion
func countGoodStrings1(low int, high int, zero int, one int) int {
	var f func(int) int
	f = func(n int) int {
		if n == 0 {
			return 1
		}
		if n < 0 {
			return 0
		}
		return (f(n-zero) + f(n-one)) % mod
	}

	count := 0
	for n := low; n <= high; n++ {
		count += f(n)
		count %= mod
	}
	return count
}
