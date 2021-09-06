package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	//最高层数
	MAX_LEVL = 16
)

//跳表节点结构体
type skipListNode struct {
	//跳表保存的值
	v interface{}
	//用于排序的分值
	score int
	//层高
	level int
	//每层前进的指针
	forwards	[]*skipListNode
}

//new skiplistnode
func newskipListNode(v interface{},score int,level int) *skipListNode {
	return &skipListNode{v:v,score: score,forwards: make([]*skipListNode, level,level),level: level}
}

//skiplist
type SkipList struct {
	//跳表头节点
	head *skipListNode
	//跳表当前层数
	level int
	//跳表长度
	length int
}

//实例化跳表对象
func NewSkipList() *SkipList {
	//头节点，便于操作
	head := newskipListNode(0,math.MaxInt32,MAX_LEVL)
	return &SkipList{head,0,1}
}

//skip length
func (sl *SkipList) Length() int {
	return sl.length
}

//skip Level
func (sl *SkipList) Level() int {
	return sl.level
}

//insert node to skiplist
func (sl *SkipList) Insert(v interface{},score int) int {
	if nil == v {
		return 1
	}

	//寻找插入位置
	cur := sl.head
	//记录每层路径
	update := [MAX_LEVL]*skipListNode{}
	i := MAX_LEVL - 1
	for ; i >=0 ;i-- {
		for nil != cur.forwards[i] {
			if cur.forwards[i].v == v {
				return 2
			}
			if cur.forwards[i].score > score {
				update[i] = cur
				break
			}
			cur = cur.forwards[i]
		} 
		if nil == cur.forwards[i]{
			update[i] = cur
		}
	}
	
	//随机算法获取该节点层数
	level := 1
	for i := 1; i < MAX_LEVL; i++ {
		if rand.Int31()%7 == 1{
			level++
		}
	}
	//创建一个新的链表节点
	newNode := newskipListNode(v,score,level)

	//原有节点连接
	for i := 0;i < level - 1; i++ {
		next := update[i].forwards[i]
		update[i].forwards[i] = newNode
		newNode.forwards[i] = next
	}

	//如果当前节点的层数大于之前跳表的层数
	//更新当前跳表层数
	if level > sl.level {
		sl.level = level
	}

	//更新跳表长度
	sl.length++

	return 0
}

