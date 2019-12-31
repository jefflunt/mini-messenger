package queue

type Fifo struct {
  messages []string   // The messages
  size uint64         // The total capacity of the queue
  used uint64         // The amount of the capacity that is currently used
  front uint64        // The index of the front of the queue
  rear uint64         // The index of the rear of the queue
}

func (f *Fifo) Alloc(size uint64) {
  f.messages = make([]string, size)
  f.size = size
  f.used = 0
  f.front = 0
  f.rear = 0
}

// Front returns the index of the front of the queue
func (f *Fifo) Front() uint64 {
  if f.Size() == 0 {
	return 0
  }

  return f.front % f.Size()
}

// Rear returns the index of the rear of the queue
func (f *Fifo) Rear() uint64 {
  if f.Size() == 0 {
	return 0
  }

  return f.rear % f.Size()
}

// Enqueue adds a message to this Fifo queue
func (f *Fifo) Enqueue(s string) {
  if f.Size() == 0 {
    return
  }

	if f.EmptyCap() == 0 {
		return // drop message, don't enqueue
  }

  f.messages[f.Rear()] = s
  f.rear++
  f.used++
}

// Dequeue removes a message from the front of this queue, returning "" if the
// queue is empty.
func (f *Fifo) Dequeue() string {
  if f.used == 0 {
    return ""
  }

  msg := f.messages[f.Front()]
  f.front++
  f.used--

  return msg
}

func (f *Fifo) Size() uint64 {
  return f.size
}

// Len returns the number of messages currently in the queue
func (f *Fifo) Len() uint64 {
  return f.used
}

// EmptyCap tells you the number of available slots for enqueuing messages. If
// the queue is full, messages that are enqueued will be dropped
func (f *Fifo) EmptyCap() uint64 {
	return f.Size() - f.Len()
}
