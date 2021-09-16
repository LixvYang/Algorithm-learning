package stack

import "fmt"


type node struct {
	next  *node
	value interface{}
}

type LinkedListStack struct {
	//top
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (l *LinkedListStack) isEmpty() bool {
	return l.topNode == nil
}

func (l *LinkedListStack) Push(v interface{}) {
	l.topNode = &node{next: l.topNode, value: v}
}

func (l *LinkedListStack) Pop() interface{} {
	if l.isEmpty() {
		return nil
	}

	v := l.topNode.value
	l.topNode = l.topNode.next
	return v
}

func (l *LinkedListStack) Top() interface{} {
	if l.isEmpty() {
		return nil
	}
	return l.topNode.value
}

func (l *LinkedListStack) Flush()  {
	l.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.isEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := this.topNode
		for nil != cur {
			fmt.Println(cur.value)
			cur = cur.next
		}
	}
}