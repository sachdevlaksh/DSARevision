package LC103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)

	if root == nil {
		return res
	}
	que := []*TreeNode{root}
    fliporder := true
	for len(que) > 0 {
		size := len(que)

		var level []int

		for i := 0; i < size; i++ {
			curr := que[0]
			que = que[1:]
			if fliporder {
				level = append(level, curr.Val)
			} else {
				level = append([]int{curr.Val}, level...)
			}

			if curr.Left != nil {
				que = append(que, curr.Left)
			}
			if curr.Right != nil {
				que = append(que, curr.Right)
			}

		}
		res = append(res, level)
        fliporder = !fliporder

	}
	return res
}
