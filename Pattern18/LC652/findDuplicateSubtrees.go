package LC652

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var dfs func(node *TreeNode) string
	dict := make(map[string][]*TreeNode)
	out := make([]*TreeNode, 0)

	dfs = func(node *TreeNode) string {
		if node == nil {
			return "N"
		}
		s := strings.Join([]string{strconv.Itoa(node.Val), dfs(node.Left), dfs(node.Right)}, ",")
		if len(dict[s]) == 1 {
			out = append(out, node)
		}
		dict[s] = append(dict[s], node)
		return s
	}
	dfs(root)
	return out
}
