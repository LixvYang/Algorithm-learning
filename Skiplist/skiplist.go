package skiplist

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	//最高层数
	MAX_LEVEL = 16
)

type skipListNode struct {
	//跳表保存的值
	v interface{}
	//对于排序的分值
	score int
	//层高
	level int
	//每层前进的指针
	forwards []*skipListNode
}

func newSkipListNode(v interface{},score int,level int) *skipListNode {
	return &skipListNode{v:v,score: score,forwards: make([]*skipListNode, level,level),level: level}
}

type SkipList struct {
	//跳表头节点
	head *skipListNode
	//跳表当前层数
	level int
	//跳表长度
	length int
}

func newSkipList() *SkipList {
	head := newSkipListNode(0,math.MinInt32,MAX_LEVEL)
	return &SkipList{head: head,level: 1,0}
}

//获取跳表长度
func (sl *SkipList) Length() int {
	return sl.length
}

//层数
func (sl *SkipList) Level() int {
	return sl.level
}


