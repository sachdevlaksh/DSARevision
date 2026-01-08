package LC572

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil{
		return false
	}

	var dfs func(node, node2 *TreeNode) bool
	dfs = func(node, node2 *TreeNode) bool {

		if node == nil && node2 == nil{
			return true
		}
		if node == nil || node2 == nil {
			return false
		}
		if node.Val != node2.Val {
			return false
		}

		return dfs(node.Left, node2.Left) && dfs(node.Right, node2.Right)
	}

	if dfs(root, subRoot){
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}
