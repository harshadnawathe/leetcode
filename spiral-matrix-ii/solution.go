package spiralmatrixii

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	x := 1
	for layer := 0; layer < (n+1)/2; layer++ {
		for i := layer; i < n-layer; i++ {
			matrix[layer][i] = x
			x++
		}

		for i := layer + 1; i < n-layer; i++ {
			matrix[i][n-layer-1] = x
			x++
		}

		for i := layer + 1; i < n-layer; i++ {
			matrix[n-layer-1][n-i-1] = x
			x++
		}

		for i := layer + 1; i < n-layer-1; i++ {
			matrix[n-i-1][layer] = x
			x++
		}
	}
	return matrix
}

func generateMatrix1(n int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	x := 1
	for c := 0; c <= n/2; c++ {
		for i := c; i < n-c; i++ {
			matrix[c][i] = x
			x++
		}

		for i := c + 1; i < n-c-1; i++ {
			matrix[i][n-c-1] = x
			x++
		}

		if c != n/2 {
			for i := n - c - 1; i >= c; i-- {
				matrix[n-c-1][i] = x
				x++
			}
		}

		for i := n - c - 2; i >= c+1; i-- {
			matrix[i][c] = x
			x++
		}
	}

	return matrix
}
