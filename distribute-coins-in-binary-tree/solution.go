package distributecoinsinbinarytree

func distributeCoins(root *TreeNode) int {
	var coinDiff func(*TreeNode) int

	numSteps := 0

	coinDiff = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		diffL := coinDiff(node.Left)
		diffR := coinDiff(node.Right)

		numSteps += abs(diffL) + abs(diffR)

		return node.Val + diffL + diffR - 1
	}

	coinDiff(root)

	return numSteps
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
