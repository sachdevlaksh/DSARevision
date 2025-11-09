package LC460

type Node struct {
	Next, Prev     *Node
	key, val, freq int
}

type DLL struct {
	head, tail *Node
}
type LFUCache struct {
	capacity int
	cache    map[int]*Node
	freqMap  map[int]*DLL
	minFreq  int
}

func newDLL() *DLL {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return &DLL{head: head, tail: tail}

}
func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		freqMap:  make(map[int]*DLL)}
}

func (this *LFUCache) Get(key int) int {
	if node, exist := this.cache[key]; exist {
		this.increaseFreq(node)
		return node.val
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.capacity == 0 {
		return
	}
	if node, exist := this.cache[key]; exist {
		node.val = value
		this.increaseFreq(node)
		return
	}

	if len(this.cache) >= this.capacity {
		list := this.freqMap[this.minFreq]
		lru := list.removeLast()
		delete(this.cache, lru.key)

	}

	newNode := &Node{key: key, val: value, freq: 1}
	this.cache[key] = newNode
	if _, ok := this.freqMap[1]; !ok {
		this.freqMap[1] = newDLL()
	}
	this.freqMap[1].addToFront(newNode)
	this.minFreq = 1
}

func (this *DLL) addToFront(node *Node) {
	node.Next = this.head.Next
	node.Prev = this.head
	this.head.Next.Prev=node
	this.head.Next=node

}
func (l *DLL) remove(node *Node) {
	node.Prev.Next= node.Next
	node.Next.Prev= node.Prev

}
func (l *DLL) removeLast() *Node {
	if l.isEmpty() {
		return nil
	}
	lastNode := l.tail.Prev
	l.remove(lastNode)
	return lastNode
}

func (this *LFUCache) increaseFreq(node *Node) {
	freq := node.freq
	this.freqMap[freq].remove(node)
	if freq == this.minFreq && this.freqMap[freq].isEmpty(){
		this.minFreq++
	}
	node.freq++
	if _, ok := this.freqMap[node.freq]; !ok{
		this.freqMap[node.freq]= newDLL()
	}
	this.freqMap[node.freq].addToFront(node)

}
func (this *DLL) isEmpty() bool {
	return this.head.Next == this.tail
}
