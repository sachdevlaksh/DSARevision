/*
LeetCode Problem #102: Binary Tree Level Order Traversal
Difficulty: Medium

Given the root of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).
*/

package LC112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0)

     if root == nil {
        return res
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
		res = append(res, level)

	}
	return res
}