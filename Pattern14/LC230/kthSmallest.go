/*
LeetCode Problem #230: Kth Smallest Element in a BST
Difficulty: Medium

Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
*/

package LC230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	curr := root
	count := 0

	for curr != nil  || len(stack) > 0{
		for curr != nil {
			stack = append(stack,curr)
			curr = curr.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		count++
		if count == k{
			return node.Val
		}

		curr = node.Right
	}
		return 0
}

