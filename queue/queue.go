package queue

type Queue interface {
  Enqueue(s string)
  Dequeue() string
  Size() uint64
  Len() uint64
	EmptyCap() uint64
}
