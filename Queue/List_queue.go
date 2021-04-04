package main

import (
    "container/list"
    "fmt"
)

type customQueue struct {
    queue *list.List
}

func (c *customQueue) Enqueue(value string) {
    c.queue.PushBack(value)
}

func (c *customQueue) Dequeue() error {
    if c.queue.Len() > 0 {
        ele := c.queue.Front()
        c.queue.Remove(ele)
    }
    return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
    if c.queue.Len() > 0 {
        if val, ok := c.queue.Front().Value.(string); ok {
            return val, nil
        }
        return "", fmt.Errorf("Peep Error: Queue Datatype is incorrect")
    }
    return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
    return c.queue.Len()
}

func main() {
    ourQueue := &customQueue{
        queue: list.New(),
    }
    ourQueue.Enqueue("老大")
    ourQueue.Enqueue("老二")
    fmt.Printf("Size: %d\n", ourQueue.Size())
	ourQueue.Enqueue("老三")
	ourQueue.Enqueue("老四")
    for ourQueue.Size() > 0 {
        frontvalue, _ := ourQueue.Front()
        fmt.Printf("Front: %s\n", frontvalue)
        fmt.Printf("Dequeue: %s\n", frontvalue)
        ourQueue.Dequeue()
    }
    fmt.Printf("Size: %d\n", ourQueue.Size())
}