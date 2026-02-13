/*
LeetCode Problem #105: Construct Binary Tree from Preorder and Inorder Traversal
Difficulty: Medium

Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
*/

package LC105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {

	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	index := 0
	var solve func(left, right int) *TreeNode
	solve = func(left, right int) *TreeNode{

		if left > right{
			return nil
		}
		rootEle := preorder[index]
		root := &TreeNode{Val: rootEle}
		index++
		rootIndex := m[rootEle]
		root.Left = solve(left,rootIndex-1)
		root.Right = solve(rootIndex+1,right)

		return root
	}
	return solve(0, len(inorder)-1)
}

func buildTreePost(inorder []int, postorder []int) *TreeNode {
	m := make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	index := len(inorder) - 1
	var solve func(left, right int) *TreeNode
	solve = func(left, right int) *TreeNode{

		if left > right{
			return nil
		}
		rootEle := postorder[index]
		root := &TreeNode{Val: rootEle}
		index--
		rootIndex := m[rootEle]
		root.Right = solve(rootIndex+1,right)
		root.Left = solve(left,rootIndex-1)

		return root
	}
	return solve(0, len(inorder)-1)
}