/*
LeetCode Problem #863: All Nodes Distance K in Binary Tree
Difficulty: Medium

Given the root of a binary tree, the value of a target node target, and an integer k, return an array of the values of all nodes that have a distance k from the target node.
*/

package LC863

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {

	if root == nil {
		return []int{}
	}
	parent := make(map[*TreeNode]*TreeNode)

	var dfs func(node, par *TreeNode)

	dfs = func(node, par *TreeNode) {
		if node == nil {
			return
		}
		parent[node] = par
		dfs(node.Left, node)
		dfs(node.Right, node)
	}
	dfs(root, nil)

	q := []*TreeNode{target}
	vis := make(map[*TreeNode]bool)
	vis[target] = true
	dis := 0

	for len(q) > 0 {
		if dis == k {
			break
		}
		size := len(q)
		for size > 0 {
			curr := q[0]
			q = q[1:]

			ne := []*TreeNode{curr.Left, curr.Right, parent[curr]}

			for _, nei := range ne {
				if nei != nil && !vis[nei] {
					vis[nei] = true
					q = append(q, nei)
				}
			}

		}
		dis++
	}
	res := make([]int, 0)
	for _, node := range q{
		res = append(res, node.Val)
	}
	return res
}
