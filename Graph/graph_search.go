package graph

import (
	"container/list"
	"fmt"
)

type Graph struct {
	adj []*list.List
	v int
}

// init graph according to capacity
func newGraph(v int) *Graph {
	graphh := &Graph{}
	graphh.v = v
	graphh.adj = make([]*list.List, v)
	for i := range graphh.adj {
		graphh.adj[i] = list.New()
	}
	return graphh
}

func (self *Graph) addEdge(s int, t int) {
	self.adj[s].PushBack(t)
	self.adj[t].PushBack(s)
}

func (self *Graph) BFS(s,t int) {
	if s == t {
		return
	}
	prev := make([]int, self.v)
	for index := range prev {
		prev[index] = -1
	}

	var queue []int
	visited := make([]bool, self.v)
	queue = append(queue, s)
	isFound := false
	for len(queue) > 0 && !isFound {
		top := queue[0]
		queue = queue[1:]
		linkedlist := self.adj[top]
		for e := linkedlist.Front(); e != nil; e = e.Next() {
			k := e.Value.(int)
			if !visited[k] {
				prev[k] = top
				if k == t {
					isFound = true
					break
 				}
				 queue = append(queue, k)
				 visited[k] = true
			}
		}
	}
	if isFound {
		printPrev(prev, s ,t)
	} else {
		fmt.Printf("no path found from %d to %d\n", s, t)
	}
}

func printPrev(prev []int, s,t int) {
	if t == s || prev[t] == -1 {
		fmt.Printf("%d", t)
	} else {
		printPrev(prev, s, prev[t])
		fmt.Print("%d", t)
	}
}

func (self *Graph) DFS(s, t int) {
	prev := make([]int, self.v)
	for i := range prev {
		prev[i] = -1
	}

	visited := make([]bool, self.v)
	visited[s] = true

	isFound := false
	self.recurse(s, t, prev, visited, isFound)
	printPrev(prev, s, t)
}

func recurse(s,t int, prev []int, visited []bool, isFound bool) {
	if isFound {
		return
	}

	visited[s] = true
	if s == t {
		isFound = true
		return
	}

	linkedlist := self.adj[s]
	for e := linkedlist.Front(); e != nil; e = e.Next() {
		k := e.Value.(int)
		if !visited[k] {
			prev[k] = s
			self.recurse(k, t, prev, visited, false)
		}
	}
}