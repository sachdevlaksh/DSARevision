/*
LeetCode Problem #1339: Maximum Product of Splitted Binary Tree
Difficulty: Medium

Given the root of a binary tree, split the binary tree into two subtrees by removing one edge such that the product of the sums of the subtrees is maximized.
*/

package LC1339

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	const mod = 1000000007
	cop := root
	totalSum := 0
	var getTotalSum func(node *TreeNode) int
	getTotalSum = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return node.Val + getTotalSum(node.Left) + getTotalSum(node.Right)
	}
	totalSum = getTotalSum(cop)

	maxProd := 0
	var getMaxProd func(node *TreeNode) int
	getMaxProd = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getMaxProd(node.Left)
		right := getMaxProd(node.Right)
		subtreeSum := node.Val + left + right
		maxProd = max(maxProd, (totalSum-subtreeSum)*subtreeSum)
		return subtreeSum
	}
	getMaxProd(root)
	return maxProd % mod
}
