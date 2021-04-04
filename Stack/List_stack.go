package main

import (
    "container/list"
    "fmt"
)

type customStack struct {
    stack *list.List
}

func (c *customStack) Push(value string) {
    c.stack.PushFront(value)
}

func (c *customStack) Pop() error {
    if c.stack.Len() > 0 {
        ele := c.stack.Front()
        c.stack.Remove(ele)
    }
    return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customStack) Front() (string, error) {
    if c.stack.Len() > 0 {
        if val, ok := c.stack.Front().Value.(string); ok {
            return val, nil
        }
        return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
    }
    return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customStack) Size() int {
    return c.stack.Len()
}

func main() {
    //声明堆栈
    ourStack := &customStack{
        stack: list.New(),
    }
    //Push 盘子A,盘子B
    ourStack.Push("盘子A")
    ourStack.Push("盘子B")
    fmt.Printf("Size:%d\n",ourStack.Size())
	ourStack.Push("盘子C")
	ourStack.Push("盘子D")
    //Pop所有盘子
    for ourStack.Size() > 0 {
        frontvalue,_ := ourStack.Front()
        fmt.Printf("Front: %s\n",frontvalue)
        fmt.Printf("Pop: %s\n",frontvalue)
        ourStack.Pop()
    }
    fmt.Printf("Size:%d\n",ourStack.Size())
}