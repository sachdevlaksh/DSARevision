package LC173

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
 Val   int
 Left  *TreeNode
 Right *TreeNode
}
type BSTIteratorOld struct {
 curr []int
}

func ConstructorOld(root *TreeNode) BSTIteratorOld {

 curr := root
 stack := make([]*TreeNode, 0)
 itr := BSTIteratorOld{}
 for curr != nil || len(stack) > 0 {
  for curr != nil {
   stack = append(stack, curr)
   curr = curr.Left
  }
  curr = stack[len(stack)-1]
  itr.curr = append(itr.curr, curr.Val)
  stack = stack[:len(stack)-1]
  curr = curr.Right
 }
 return itr
}

func (this *BSTIteratorOld) Next() int {
 if len(this.curr) > 0 {
  out := this.curr[0]
  this.curr = this.curr[1:]
  return out

 }
 return -1
}

func (this *BSTIteratorOld) HasNext() bool {

 if len(this.curr) > 0 {
  return true
 }
 return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

type BSTIterator struct {
 stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
 itr := BSTIterator{}
 itr.CollectLeft(root)
 return itr
}

func (this *BSTIterator) CollectLeft(node *TreeNode) {

 curr := node
 for curr != nil {
  this.stack = append(this.stack, curr)
  curr = curr.Left
 }

}
func (this *BSTIterator) Next() int {
 curr := this.stack[len(this.stack)-1]
 this.stack = this.stack[:len(this.stack)-1]
 if curr.Right != nil {
   this.CollectLeft(curr.Right)
 }
 return curr.Val
}

func (this *BSTIterator) HasNext() bool {

 return len(this.stack) > 0
}
