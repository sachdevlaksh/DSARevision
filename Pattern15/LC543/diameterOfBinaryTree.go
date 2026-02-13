/*
LeetCode Problem #543: Diameter of Binary Tree
Difficulty: Easy

Given the root of a binary tree, return the length of the diameter of the tree. The diameter of a binary tree is the length of the longest path between any two nodes in a tree.
*/

package LC543

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {

	var getHeight func(curr *TreeNode) int
	diameter := 0

	getHeight = func(curr *TreeNode) int {

		if curr == nil {
			return 0
		}
		lh := getHeight(curr.Left)
		rh := getHeight(curr.Right)

		diameter = max(diameter, lh+rh)

		return 1 + max(lh, rh)
	}
	getHeight(root)
	return diameter
}
