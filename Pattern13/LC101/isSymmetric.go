/*
LeetCode Problem #101: Symmetric Tree
Difficulty: Easy

Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
*/

package LC101

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
        return true
    }

	return solve(root.Left, root.Right)
}

func solve(t1,t2 *TreeNode) bool{
	if t1 == nil && t2 == nil{
		return true
	}
	if t1 == nil || t2 == nil{
		return false
	}

	return (t1.Val==t2.Val ) && solve(t1.Right, t2.Left) && solve(t1.Left, t2.Right)

}