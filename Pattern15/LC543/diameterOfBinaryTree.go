package LC543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	var getHeight func(curr *TreeNode) int
	diameter := 0

	getHeight = func(curr *TreeNode) int {

		if curr == nil {
			return 0
		}
		lh := getHeight(curr.Left)
		rh := getHeight(curr.Right)

		diameter = max(diameter, lh+rh)

		return 1 + max(lh, rh)
	}
	getHeight(root)
	return diameter
}
