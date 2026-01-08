package LC104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var dfs func(curr *TreeNode) int
	dfs = func(curr *TreeNode) int {
		if curr == nil {
			return 0
		}
		return 1 + max(dfs(curr.Left), dfs(curr.Right))
	}

	return dfs(root)
}
func maxDepthItr(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	height := 0
	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			curr := q[0]
			q = q[1:]

			if curr.Left != nil {
				q = append(q, curr.Left)
			}
			if curr.Right != nil {
				q = append(q, curr.Right)
			}
			size--
		}
		height++
	}
	return height
}
