package leetcode

type LRUCache struct {
	capacity int
	m map[int]*Node
	head, tail *Node
}

type Node struct {
	Key int
	Value int
	Pre, Next *Node
}

func (this *LRUCache) Get(key int) int {
	if v, ok := this.m[key]; ok {
			this.moveToHead(v)
			return v.Value
	}
	return -1
}

func (this *LRUCache) moveToHead(node *Node) {
	this.deleteNode(node)
	this.addToHead(node)
}

func (this *LRUCache) deleteNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) removeTail() int {
	node := this.tail.Pre
	this.deleteNode(node)
	return node.Key
}

func (this *LRUCache) addToHead(node *Node) {
	this.head.Next.Pre = node
	node.Next = this.head.Next
	node.Pre = this.head
	this.head.Next = node
}

func (this *LRUCache) Put(key int, value int)  {
	if v, ok := this.m[key]; ok {
			v.Value = value
			this.moveToHead(v)
			return
	}

	if this.capacity == len(this.m) {
			rmKey := this.removeTail()
			delete(this.m, rmKey)
	}

	newNode := &Node{Key: key, Value: value}
	this.addToHead(newNode)
	this.m[key] = newNode
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head
	return LRUCache{
			capacity: capacity, 
			m: map[int]*Node{}, 
			head: head, 
			tail: tail,
	}
}