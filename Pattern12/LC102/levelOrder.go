package LC112

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0)

     if root == nil {
        return res
    }
	que := []*TreeNode{root}

	for len(que) > 0 {
		size := len(que)
		var level []int

		for i := 0; i < size; i++ {
			curr := que[0]
			que = que[1:]
			level = append(level, curr.Val)

			if curr.Left != nil {
				que = append(que, curr.Left)
			}
			if curr.Right != nil {
				que = append(que, curr.Right)
			}
		}
		res = append(res, level)

	}
	return res
}