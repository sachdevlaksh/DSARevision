package LC297

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {

	var res []string
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			res = append(res, "#")
			return
		}
		res = append(res, strconv.Itoa(node.Val))
		dfs(node.Left)
		dfs(node.Right)

	}
	dfs(root)
	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
        return nil
    }
	nodes := strings.Split(data, ",")
	var i int

	var build func() *TreeNode

	build = func() *TreeNode {
		if nodes[i] == "#" {
			i++
			return nil
		}
		val, _ := strconv.Atoi(nodes[i])
		i++
		node := &TreeNode{Val: val}
		node.Left = build()
		node.Right = build()
		return node
	}

	return build()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
