package maximumwidthofbinarytree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := 0
	level := []*indexedTreeNode{{root, 0}}
	for len(level) > 0 {
		if w := width(level); max < w {
			max = w
		}

		nextLevel := []*indexedTreeNode{}
		for _, node := range level {
			if leftChild := node.LeftChild(); leftChild != nil {
				nextLevel = append(nextLevel, leftChild)
			}
			if rightChild := node.RightChild(); rightChild != nil {
				nextLevel = append(nextLevel, rightChild)
			}
		}
		level = nextLevel
	}
	return max
}

type indexedTreeNode struct {
	*TreeNode
	index int
}

func (node *indexedTreeNode) LeftChild() *indexedTreeNode {
	if node.Left == nil {
		return nil
	}
	return &indexedTreeNode{node.Left, node.index * 2}
}

func (node *indexedTreeNode) RightChild() *indexedTreeNode {
	if node.Right == nil {
		return nil
	}
	return &indexedTreeNode{node.Right, node.index*2 + 1}
}

func width(a []*indexedTreeNode) int {
	if len(a) == 0 {
		return 0
	}

	first, last := 0, len(a)-1
	return 1 + a[last].index - a[first].index
}
