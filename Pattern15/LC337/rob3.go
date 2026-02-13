/*
LeetCode Problem #337: House Robber III
Difficulty: Medium

The thief has found himself a new place for his thievery again. There is only one entrance to this area, called root. Besides the root, each house has one parent house. After a tour, the smart thief realized that all houses in this place form a binary tree.
*/

package LC337

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {

	var dfs func(root *TreeNode) (int, int)

	dfs = func(curr *TreeNode) (int, int) {

		if curr == nil{
			return 0,0
		}

		robLeftskip, robLeftTaken := dfs(curr.Left)
		robRightskip, robRightTaken := dfs(curr.Right)

		taken := curr.Val + robLeftskip + robRightskip

		skip := max(robLeftTaken, robLeftskip) + max(robRightskip, robRightTaken)

		return skip, taken
	}

	rootSkip, rootTaken := dfs(root)

	return max(rootSkip, rootTaken)

}
