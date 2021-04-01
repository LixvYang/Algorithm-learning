package Stack

import (
    "fmt"
    "sync"
)

type customQueue struct {
    stack []string
    lock  sync.RWMutex
}

func (c *customQueue) Push(name string) {
    c.lock.Lock()
    defer c.lock.Unlock()
    c.stack = append(c.stack, name)
}

func (c *customQueue) Pop() error {
    len := len(c.stack)
    if len > 0 {
        c.lock.Lock()
        defer c.lock.Unlock()
        c.stack = c.stack[:len-1]
        return nil
    }
    return fmt.Errorf("Pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
    len := len(c.stack)
    if len > 0 {
        c.lock.Lock()
        defer c.lock.Unlock()
        return c.stack[len-1], nil
    }
    return "", fmt.Errorf("Peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
    return len(c.stack)
}

func (c *customQueue) Empty() bool {
    return len(c.stack) == 0
}

func main() {
    customQueue := &customQueue{
        stack: make([]string, 0),
    }
    fmt.Printf("Push: A\n")
    customQueue.Push("A")
    fmt.Printf("Push: B\n")
    customQueue.Push("B")
    fmt.Printf("Size: %d\n", customQueue.Size())
    for customQueue.Size() > 0 {
        frontVal, _ := customQueue.Front()
        fmt.Printf("Front: %s\n", frontVal)
        fmt.Printf("Pop: %s\n", frontVal)
        customQueue.Pop()
    }
    fmt.Printf("Size: %d\n", customQueue.Size())
}