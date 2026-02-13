/*
LeetCode Problem #98: Validate Binary Search Tree
Difficulty: Medium

Given the root of a binary tree, determine if it is a valid binary search tree (BST).
*/

package LC98

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func isValidBSTRec(root *TreeNode) bool {
    var dfs func(curr *TreeNode, min, max *int) bool
    dfs = func(curr *TreeNode, min, max *int) bool {
        if curr == nil {
            return true
        }
        if min != nil && curr.Val < *min {
            return false
        }
        if max != nil && curr.Val > *max {
            return false
        }
        return dfs(curr.Left, min, &curr.Val) && dfs(curr.Right, &curr.Val, max)
    }
    return dfs(root, nil, nil)
}
func isValidBST(root *TreeNode) bool {
    stacks := make([]*TreeNode, 0)
    curr := root
    var prev *int
    for curr != nil || len(stacks) > 0 {
        for curr != nil {
            stacks = append(stacks, curr)
            curr = curr.Left
        }
        curr = stacks[len(stacks)-1]
        stacks = stacks[:len(stacks)-1]
        if prev != nil && curr.Val <= *prev {
            return false
        }
        prev = &curr.Val
        curr = curr.Right

    }
    return true
}
