package LC124

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	maxi := math.MinInt
	calculateSum(root,&maxi)
	return maxi
}

func calculateSum(node *TreeNode, maxi *int) int{
	if node == nil{
		return 0
	}

	maxleft := max(0, calculateSum(node.Left, maxi))
	maxRight := max(0, calculateSum(node.Right, maxi))

	*maxi = max(*maxi, node.Val+maxleft+maxRight)
	return max(maxleft, maxRight) + node.Val

}
