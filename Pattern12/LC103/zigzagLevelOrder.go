/*
LeetCode Problem #103: Binary Tree Zigzag Level Order Traversal
Difficulty: Medium

Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. (i.e., from left to right, then right to left for the next level and alternate between).
*/

package LC103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)

	if root == nil {
		return res
	}
	que := []*TreeNode{root}
    fliporder := true
	for len(que) > 0 {
		size := len(que)

		var level []int

		for i := 0; i < size; i++ {
			curr := que[0]
			que = que[1:]
			if fliporder {
				level = append(level, curr.Val)
			} else {
				level = append([]int{curr.Val}, level...)
			}

			if curr.Left != nil {
				que = append(que, curr.Left)
			}
			if curr.Right != nil {
				que = append(que, curr.Right)
			}

		}
		res = append(res, level)
        fliporder = !fliporder

	}
	return res
}
