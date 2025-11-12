package LC100

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var sb strings.Builder
	var sb1 strings.Builder
	itrTree(p, &sb)
	itrTree(q, &sb1)
	if sb.String() == sb1.String(){
		return true
	}
	return false
}

func itrTree(p *TreeNode, sb *strings.Builder) {

	if p != nil {
		sb.WriteString(strconv.Itoa(p.Val))
		if p.Left != nil {
			itrTree(p.Left, sb)
		} else {
			sb.WriteString("#")
		}
		if p.Right != nil {
			itrTree(p.Right, sb)
		} else {
			sb.WriteString("#")
		}
	}
}
