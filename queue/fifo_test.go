package queue

import "testing"

func TestFifo_Alloc_Size(t *testing.T) {
  fifo0 := &Fifo{}

  fifo3 := &Fifo{}
  fifo3.Alloc(3)

  testCases := []struct {
    name string
    queue *Fifo
    want uint64
  }{
    {
      name: "0",
      queue: fifo0,
      want: 0,
    },{
      name: "3",
      queue: fifo3,
      want: 3,
    },
  }

  for _, testCase := range testCases {
    t.Run(testCase.name, func(t *testing.T) {
      got := testCase.queue.Size()
      if testCase.want != got {
        t.Errorf("Alloc()/Size(): want %d, got %d", testCase.want, got)
      }
    })
  }
}

func TestFifo_Front(t *testing.T) {
  fifo0Enq0 := &Fifo{}

  fifo3Enq1 := &Fifo{}
	fifo3Enq1.Alloc(3)
  fifo3Enq1.Enqueue("msg1")

  fifo3Enq2Deq1 := &Fifo{}
	fifo3Enq2Deq1.Alloc(10)
  fifo3Enq2Deq1.Enqueue("msg1")
  fifo3Enq2Deq1.Enqueue("msg2")
  fifo3Enq2Deq1.Dequeue()

  fifo3Enq3Deq2 := &Fifo{}
	fifo3Enq3Deq2.Alloc(10)
  fifo3Enq3Deq2.Enqueue("msg1")
  fifo3Enq3Deq2.Enqueue("msg2")
  fifo3Enq3Deq2.Dequeue()
  fifo3Enq3Deq2.Dequeue()

  fifo3Enq3Deq4 := &Fifo{}
  fifo3Enq3Deq4.Alloc(10)
  fifo3Enq3Deq4.Enqueue("msg1")
  fifo3Enq3Deq4.Enqueue("msg2")
  fifo3Enq3Deq4.Dequeue()
  fifo3Enq3Deq4.Dequeue()
  fifo3Enq3Deq4.Dequeue()

	fifo3Enq3Deq3Enq3Deq1 := &Fifo{}
	fifo3Enq3Deq3Enq3Deq1.Alloc(3)
  fifo3Enq3Deq3Enq3Deq1.Enqueue("msg1")
  fifo3Enq3Deq3Enq3Deq1.Enqueue("msg2")
  fifo3Enq3Deq3Enq3Deq1.Enqueue("msg3")
  fifo3Enq3Deq3Enq3Deq1.Dequeue()
  fifo3Enq3Deq3Enq3Deq1.Dequeue()
  fifo3Enq3Deq3Enq3Deq1.Dequeue()
  fifo3Enq3Deq3Enq3Deq1.Enqueue("msg4")
  fifo3Enq3Deq3Enq3Deq1.Enqueue("msg5")
  fifo3Enq3Deq3Enq3Deq1.Enqueue("msg6")
  fifo3Enq3Deq3Enq3Deq1.Dequeue()

  testCases := []struct {
    name string
    queue *Fifo
    want uint64
  }{
    {
      name: "0",
      queue: fifo0Enq0,
      want: 0,
    },{
      name: "3, enq 1",
      queue: fifo3Enq1,
      want: 0,
    },{
      name: "0, enq 2, deq 1",
      queue: fifo3Enq2Deq1,
      want: 1,
    },{
      name: "0, enq 3, deq 2",
      queue: fifo3Enq3Deq2,
      want: 2,
    },{
      name: "3, enq 3, deq 4",
      queue: fifo3Enq3Deq4,
      want: 2,
    },{
			name: "3, enq 3, deq 3, enq 3, deq 1",
			queue: fifo3Enq3Deq3Enq3Deq1,
      want: 1,
		},
  }

  for _, testCase := range testCases {
    t.Run(testCase.name, func(t *testing.T) {
      got := testCase.queue.Front()
      if testCase.want != got {
        t.Errorf("Front(): want %d, got %d", testCase.want, got)
      }
    })
  }
}

func TestFifo_Rear(t *testing.T) {
  fifo0Enq0 := &Fifo{}
	fifo0Enq0.Alloc(3)

  fifo3Enq1 := &Fifo{}
	fifo3Enq1.Alloc(3)
  fifo3Enq1.Enqueue("msg1")

  fifo3Enq2 := &Fifo{}
	fifo3Enq2.Alloc(3)
  fifo3Enq2.Enqueue("msg1")
  fifo3Enq2.Enqueue("msg2")

  fifo3Enq4 := &Fifo{}
  fifo3Enq4.Alloc(3)
  fifo3Enq4.Enqueue("msg1")
  fifo3Enq4.Enqueue("msg2")
  fifo3Enq4.Enqueue("msg3")
  fifo3Enq4.Enqueue("msg4")

  fifo5Enq4 := &Fifo{}
  fifo5Enq4.Alloc(5)
  fifo5Enq4.Enqueue("msg1")
  fifo5Enq4.Enqueue("msg2")
  fifo5Enq4.Enqueue("msg3")
  fifo5Enq4.Enqueue("msg4")

  testCases := []struct {
    name string
    queue *Fifo
    want uint64
  }{
    {
      name: "0, enq 0",
      queue: fifo0Enq0,
      want: 0,
    },{
      name: "3, enq 1",
      queue: fifo3Enq1,
      want: 1,
    },{
      name: "3, enq 2",
      queue: fifo3Enq2,
      want: 2,
    },{
      name: "3, enq 4",
      queue: fifo3Enq4,
      want: 0,
		},{
      name: "3, enq 2",
      queue: fifo5Enq4,
      want: 4,
		},
  }

  for _, testCase := range testCases {
    t.Run(testCase.name, func(t *testing.T) {
      got := testCase.queue.Rear()
      if testCase.want != got {
        t.Errorf("Rear(): want %d, got %d", testCase.want, got)
      }
    })
  }
}

func TestFifo_Len(t *testing.T) {
  fifoEnq0Alloc0 := &Fifo{}

	fifoEnq0Alloc3 := &Fifo{}
  fifoEnq0Alloc3.Alloc(3)

  fifoEnq1Alloc3 := &Fifo{}
  fifoEnq1Alloc3.Alloc(3)
  fifoEnq1Alloc3.Enqueue("msg1")

  fifoEnq3Alloc3 := &Fifo{}
  fifoEnq3Alloc3.Alloc(3)
  fifoEnq3Alloc3.Enqueue("msg1")
  fifoEnq3Alloc3.Enqueue("msg2")
  fifoEnq3Alloc3.Enqueue("msg3")

  testCases := []struct {
    name string
    queue *Fifo
    want uint64
  }{
		{
			name: "0, enq 0",
			queue: fifoEnq0Alloc0,
			want: 0,
		},{
			name: "3, enq 0",
			queue: fifoEnq0Alloc3,
			want: 0,
		},{
			name: "3, enq 1",
			queue: fifoEnq1Alloc3,
			want: 1,
		},{
			name: "3, enq 3",
			queue: fifoEnq3Alloc3,
			want: 3,
		},
  }

  for _, testCase := range testCases {
    t.Run(testCase.name, func(t *testing.T) {
      got := testCase.queue.Len()
      if testCase.want != got {
        t.Errorf("Len(): want %d, got %d", testCase.want, got)
      }
    })
  }
}

func TestFifo_EmptyCap(t *testing.T) {
  fifo0Enq0 := &Fifo{}

	fifo3Enq0 := &Fifo{}
  fifo3Enq0.Alloc(3)

  fifo3Enq1 := &Fifo{}
  fifo3Enq1.Alloc(3)
  fifo3Enq1.Enqueue("msg1")

  fifo3Enq3 := &Fifo{}
  fifo3Enq3.Alloc(3)
  fifo3Enq3.Enqueue("msg1")
  fifo3Enq3.Enqueue("msg2")
  fifo3Enq3.Enqueue("msg3")

  fifo3Enq3Deq4 := &Fifo{}
  fifo3Enq3Deq4.Alloc(3)
  fifo3Enq3Deq4.Enqueue("msg1")
  fifo3Enq3Deq4.Enqueue("msg2")
  fifo3Enq3Deq4.Enqueue("msg3")
  fifo3Enq3Deq4.Dequeue()
  fifo3Enq3Deq4.Dequeue()
  fifo3Enq3Deq4.Dequeue()
  fifo3Enq3Deq4.Dequeue()

  testCases := []struct {
    name string
    queue *Fifo
    want uint64
  }{
		{
			name: "0, enq 0",
			queue: fifo0Enq0,
			want: 0,
		},{
			name: "3, enq 0",
			queue: fifo3Enq0,
			want: 3,
		},{
			name: "3, enq 1",
			queue: fifo3Enq1,
			want: 2,
		},{
			name: "3, enq 3",
			queue: fifo3Enq3,
			want: 0,
		},{
      name: "3, enq 3, deq 4",
      queue: fifo3Enq3Deq4,
      want: 3,
    },
  }

  for _, testCase := range testCases {
    t.Run(testCase.name, func(t *testing.T) {
      got := testCase.queue.EmptyCap()
      if testCase.want != got {
        t.Errorf("EmptyCap(): want %d, got %d", testCase.want, got)
      }
    })
  }
}

func TestFifo_EnqueueDequeue(t *testing.T) {
  fifo0Enq0 := &Fifo{}

	fifo3Enq0 := &Fifo{}
  fifo3Enq0.Alloc(3)

	fifo3Enq1 := &Fifo{}
  fifo3Enq1.Alloc(3)
  fifo3Enq1.Enqueue("msg1")

  fifo3Enq3Deq1 := &Fifo{}
  fifo3Enq3Deq1.Alloc(3)
  fifo3Enq3Deq1.Enqueue("msg1")
  fifo3Enq3Deq1.Enqueue("msg2")
  fifo3Enq3Deq1.Enqueue("msg3")
  fifo3Enq3Deq1.Dequeue()

  fifo3Enq3Deq2 := &Fifo{}
  fifo3Enq3Deq2.Alloc(3)
  fifo3Enq3Deq2.Enqueue("msg1")
  fifo3Enq3Deq2.Enqueue("msg2")
  fifo3Enq3Deq2.Enqueue("msg3")
  fifo3Enq3Deq2.Dequeue()
  fifo3Enq3Deq2.Dequeue()

  fifo3Enq3Deq4 := &Fifo{}
  fifo3Enq3Deq4.Alloc(3)
  fifo3Enq3Deq4.Enqueue("msg1")
  fifo3Enq3Deq4.Enqueue("msg2")
  fifo3Enq3Deq4.Enqueue("msg3")
  fifo3Enq3Deq4.Dequeue()
  fifo3Enq3Deq4.Dequeue()
  fifo3Enq3Deq4.Dequeue()

  testCases := []struct {
    name string
    queue *Fifo
    want string
  }{
		{
			name: "0, enq 0",
			queue: fifo0Enq0,
			want: "",
		},{
			name: "3, enq 0",
			queue: fifo3Enq0,
			want: "",
		},{
			name: "3, enq 1",
			queue: fifo3Enq1,
			want: "msg1",
		},{
			name: "3, enq 3, deq 1",
			queue: fifo3Enq3Deq1,
			want: "msg2",
		},{
			name: "3, enq 3, deq 2",
			queue: fifo3Enq3Deq2,
			want: "msg3",
		},{
			name: "3, enq 3, deq 4",
			queue: fifo3Enq3Deq4,
			want: "",
		},
  }

  for _, testCase := range testCases {
    t.Run(testCase.name, func(t *testing.T) {
      got := testCase.queue.Dequeue()
      if testCase.want != got {
        t.Errorf("Enqueue()/Dequeue(): want %q, got %q", testCase.want, got)
      }
    })
  }
}
