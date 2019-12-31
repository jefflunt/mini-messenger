package main

import (
  "fmt"
  "github.com/jefflunt/mini-messenger/queue"
)

func main() {
  fifo := queue.Fifo{}
  fifo.Enqueue("foo")
  fifo.Enqueue("bar")
  fifo.Enqueue("baz")

  fmt.Println(fifo.Dequeue())
  fmt.Println(fifo.Dequeue())
  fmt.Println(fifo.Dequeue())
  fmt.Println(fifo.Dequeue())
}
