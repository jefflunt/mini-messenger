package main

import (
  "fmt"
  "github.com/jefflunt/mini-messenger/queue"
)

func main() {
  fifo := queue.Fifo{}
  fifo.Alloc(3)
  fifo.Enqueue("foo")
  fifo.Enqueue("bar")
  fifo.Enqueue("baz")

  fmt.Printf("%q\n", fifo.Dequeue())
  fmt.Printf("%q\n", fifo.Dequeue())
  fmt.Printf("%q\n", fifo.Dequeue())
  fmt.Printf("%q\n", fifo.Dequeue())
}
