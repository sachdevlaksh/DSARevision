/*
LeetCode Problem #199: Binary Tree Right Side View
Difficulty: Medium

Given the root of a binary tree, imagine yourself standing on the right side of it, return the values of the nodes you can see ordered from top to bottom.
*/

package LC199

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {

	res := make([]int, 0)

	if root == nil {
		return []int{}
	}
	que := []*TreeNode{root}

	for len(que) > 0 {
		size := len(que)
		var level []int

		for i := 0; i < size; i++ {
			curr := que[0]
			que = que[1:]
			level = append(level, curr.Val)

			if curr.Left != nil {
				que = append(que, curr.Left)
			}
			if curr.Right != nil {
				que = append(que, curr.Right)
			}
		}
		res = append(res, level[len(level)-1])

	}
	return res
}
