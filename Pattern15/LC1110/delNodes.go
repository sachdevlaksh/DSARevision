package LC1110

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {

	hm := make(map[int]bool)
	out := []*TreeNode{}

	for _, v := range to_delete {
		hm[v] = true
	}

	var dfs func(node *TreeNode, parentDeleteed bool) *TreeNode
	dfs = func(node *TreeNode, parentDeleteed bool) *TreeNode {

		if node == nil {
			return nil
		}

		deleted := hm[node.Val]
		if !deleted && parentDeleteed {
			out = append(out, node)
		}

		node.Left = dfs(node.Left, deleted)
		node.Right = dfs(node.Right, deleted)

		if deleted {
			return nil
		}
		return node
	}

	dfs(root, true)
	return out
}
