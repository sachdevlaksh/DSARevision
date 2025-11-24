package LC110

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}


func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	l := getHeight(root.Left)

	r := getHeight(root.Right)

	if math.Abs(float64(l-r)) > 1 {
		return false
	}

	ln := isBalanced(root.Left)
	lr := isBalanced(root.Right)

    if !ln || !lr{
        return false
    }

	return true
}

func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(getHeight(node.Left), getHeight(node.Right))
}

func isBalancedOp(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return -1 != getHeightOp(root)
}

func getHeightOp(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lh := getHeightOp(node.Left)
	if -1 == lh{
		return -1
	}
	rh := getHeightOp(node.Right)
	if -1 == rh{
		return -1
	}
	if math.Abs(float64(lh-rh)) > 1 {
		return -1
	}

	return 1 + max(lh, rh)
}
