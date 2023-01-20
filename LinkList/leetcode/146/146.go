package leetcode

type LRUCache struct {
    capacity int
    m map[int]*Node
    head, tail *Node
}

type Node struct {
    next, pre *Node
    Key, Val int
}

func Constructor(capacity int) LRUCache {
    head, tail := &Node{}, &Node{}
    head.next = tail
    tail.pre = head
    return LRUCache{
        capacity: capacity,
        m: map[int]*Node{},
        head: head,
        tail: tail,
    }
}

func (this*LRUCache) deleteNode(node *Node) {
    node.pre.next = node.next
    node.next.pre = node.pre
}

func (this *LRUCache) moveToHead(node *Node) {
    this.deleteNode(node)
    this.addToHead(node)
}

func (this *LRUCache) addToHead(node *Node) {
    node.next = this.head.next
    this.head.next.pre = node
    this.head.next = node
    node.pre = this.head
}


func (this *LRUCache) Get(key int) int {
    if node, ok := this.m[key]; ok {
        this.moveToHead(node)
        return node.Val
    }
    return -1
}

func (this *LRUCache) deleteTail() int {
    node := this.tail.pre
    this.deleteNode(node)
    return node.Key
}


func (this *LRUCache) Put(key int, value int)  {
    if node, ok := this.m[key]; ok {
        node.Val = value
        this.moveToHead(node)
        return
    }
    if this.capacity == len(this.m) {
        rmKey := this.deleteTail()
        delete(this.m, rmKey)
    }
    newNode := &Node{Key: key, Val: value}
    this.addToHead(newNode)
    this.m[key] = newNode
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */