package LC146

type Node struct {
	key, value int
	Next, Prev *Node
}
type LRUCache struct {
	cache map[int]*Node
	capacity int
	head *Node
	tail *Node

}


func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head

	return LRUCache{
		cache: make(map[int]*Node),
		head: head,
		tail: tail,
		capacity: capacity,
	}


}


func (this *LRUCache) Get(key int) int {

	if node, ok := this.cache[key]; ok{
		this.moveToFront(node)
		return node.value
	}

	return -1

}


func (this *LRUCache) Put(key int, value int)  {

	if node , exist := this.cache[key]; exist{
		this.remove(node)
		node.value = value
		this.moveToFront(node)
		return
	}

	newNode := &Node{key: key,value: value}
	this.cache[key] = newNode
	this.add(newNode)

	if len(this.cache) > this.capacity{
		lru := this.tail.Prev
		this.remove(lru)
		delete(this.cache, lru.key)

	}


}



func (this *LRUCache) moveToFront(node *Node)  {
	this.remove(node)
	this.add(node)

}

func (this *LRUCache) remove(node *Node)  {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev= prev
}
func (this *LRUCache) add(node *Node)  {

	node.Prev= this.head
	node.Next= this.head.Next
	this.head.Next.Prev = node
	this.head.Next = node
}