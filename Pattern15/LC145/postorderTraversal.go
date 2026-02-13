/*
LeetCode Problem #145: Binary Tree Postorder Traversal
Difficulty: Easy

Given the root of a binary tree, return the postorder traversal of its nodes' values.
*/

package LC145

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {

	if root == nil {
		return []int{}
	}
	out := make([]int, 0)
	stack := []*TreeNode{}
	curr := root
	var lastVisted *TreeNode
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		peak := stack[len(stack)-1]
		if peak.Right != nil && lastVisted != peak.Right {
			curr = peak.Right
		} else {
			out = append(out, peak.Val)
			lastVisted = peak
			stack = stack[:len(stack)-1]
			curr = nil
		}

	}

	return out
}
