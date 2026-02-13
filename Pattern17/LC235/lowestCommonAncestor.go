/*
LeetCode Problem #235: Lowest Common Ancestor of a Binary Search Tree
Difficulty: Easy

Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.
*/

package LC235

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func maximizeExpressionOfThree(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]<nums[j]
	})

	n := len(nums)
	if n < 3{
		return 0
	}
	return nums[n-1] + nums[n-2] - nums[0]
}

func MinLengthAfterRemovals(s string) int {

	li1 := make([]int, 0)
	li2 := make([]int, 0)
	li1 = append(li1, 0)

	j := 1
	for j < len(s) && s[0] == s[j] {
		j++
	}
	if j >= len(s)-1{
		return len(s)
	}else{
		li2 = append(li2, j+1)
	}

	for i, si := range s{
		n1 := len(li1)
		n2 := len(li2)
		if si != int32(s[li1[n1-1]]) && li1[n1-1] > i{
			li2 = append(li2, i)
		}
		if si != int32(s[li2[n2-1]]) && li1[n2-1] > i{
			li1 = append(li1, i)
		}
	}
	fmt.Println(li1)
	fmt.Println(li2)

	return 0
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	if root == nil{
		return nil
	}

	if root.Val == p.Val || root.Val == q.Val{
		return root

	}
	lv := lowestCommonAncestor(root.Left, p,q)
	rv := lowestCommonAncestor(root.Right, p,q)

	if lv != nil && rv != nil{
		return root
	}else if lv != nil {
		return lv
	}else if rv != nil{
		return rv
	}else{
		return nil
	}

}

