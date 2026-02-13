/*
LeetCode Problem #94: Binary Tree Inorder Traversal
Difficulty: Easy

Given the root of a binary tree, return the inorder traversal of its nodes' values.
*/

package LC94

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalRec(root *TreeNode) []int {
	res := make([]int, 0)
	var dfs func(curr *TreeNode)
	dfs = func(curr *TreeNode) {
		if curr == nil {
			return
		}
		if curr.Left != nil {
			dfs(curr.Left)
		}
		res = append(res, curr.Val)
		if curr.Right != nil {
			dfs(curr.Right)
		}
		return
	}
	dfs(root)
	return res
}

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	stacks := make([]*TreeNode, 0)
	curr := root
	for curr != nil || len(stacks) > 0 {
		for curr != nil {
			stacks = append(stacks, curr)
			curr = curr.Left
		}
		curr = stacks[len(stacks)-1]
		res = append(res, curr.Val)
		stacks = stacks[:len(stacks)-1]
		curr = curr.Right
	}
	return res
}
