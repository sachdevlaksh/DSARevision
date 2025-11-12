package LC114

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	prev := &TreeNode{}
	prev = nil

	var flat func(node *TreeNode)
	flat = func(node *TreeNode) {
		if node == nil {
			return
		}
		flat(node.Left)
		flat(node.Right)

		node.Right = prev
		node.Left = nil
		prev = node
	}

	flat(root)

}

func flattenStack(root *TreeNode) {
    if root == nil {
        return
    }

    stack := []*TreeNode{root}

    for len(stack) > 0 {
        // pop
        n := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        // push right then left (so left is processed first)
        if n.Right != nil {
            stack = append(stack, n.Right)
        }
        if n.Left != nil {
            stack = append(stack, n.Left)
        }

        // link current node's right to next node in preorder (top of stack)
        n.Left = nil
        if len(stack) > 0 {
            n.Right = stack[len(stack)-1]
        } else {
            n.Right = nil
        }
    }
}
func flattenMorris(root *TreeNode) {
    curr := root
    for curr != nil {
        if curr.Left != nil {
            // find rightmost of left subtree
            pre := curr.Left
            for pre.Right != nil {
                pre = pre.Right
            }
            // connect original right subtree after the predecessor
			pre.Right = curr.Right
            // move left subtree to right
            curr.Right = curr.Left
            curr.Left = nil
        }
        // advance
        curr = curr.Right
}
